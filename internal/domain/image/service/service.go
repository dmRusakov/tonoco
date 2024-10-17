package service

import (
	"context"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/dmRusakov/tonoco/internal/domain/image/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/jpeg"
	"net/http"
	"os"
	"time"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type Item = entity.Image
type Filter = entity.ImageFilter

type Repository interface {
	Get(context.Context, *Filter) (*Item, error)
	List(context.Context, *Filter) (*map[uuid.UUID]Item, error)
	Create(context.Context, *Item) (*uuid.UUID, error)
	Update(context.Context, *Item) error
	Patch(context.Context, *uuid.UUID, *map[string]interface{}) error
	UpdatedAt(context.Context, *uuid.UUID) (*time.Time, error)
	TableIndexCount(context.Context) (*uint64, error)
	MaxSortOrder(context.Context) (*uint64, error)
	Delete(context.Context, *uuid.UUID) error
}

type Service struct {
	repository        model.Storage
	originalFilesPath string
	webFilesPath      string
	feedFilesPath     string
	sizes             map[string]uint
}

func NewService(repository *model.Model) *Service {
	return &Service{
		repository:        repository,
		originalFilesPath: "assets/images/original/",
		webFilesPath:      "assets/images/web/",
		feedFilesPath:     "assets/images/feed/",
		sizes: map[string]uint{
			"full_path":   2000,
			"large_path":  1500,
			"medium_path": 800,
			"grid_path":   400,
			"thumb_path":  200,
		},
	}
}

func (s *Service) Compression(ctx context.Context, param *entity.ImageCompression) error {
	// Get the image from the repository
	img, err := s.Get(ctx, &entity.ImageFilter{
		Ids: param.Ids,
	})
	if err != nil {
		return err
	}

	// Open the original image file
	originalFilePath := s.originalFilesPath + img.OriginPath
	file, err := os.Open(originalFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// get file format (jpg, png, etc)
	format, err := func(file *os.File) (string, error) {
		// Read the first 512 bytes of the file
		buffer := make([]byte, 512)
		_, err := file.Read(buffer)
		if err != nil {
			return "", err
		}

		// Reset the file pointer
		_, err = file.Seek(0, 0)
		if err != nil {
			return "", err
		}

		// Get the file format
		format := http.DetectContentType(buffer)
		return format, nil
	}(file)
	if err != nil {
		return err
	}

	switch format {
	case "image/jpeg":

	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	return nil
}

func (s *Service) compressJpg(
	file os.File,
	filename string,
	filePath string,
	size uint,
	compression uint,
	watermarkText string,
	watermarkFont string,
) (string, string, error) {
	srcImage, err := jpeg.Decode(&file)
	if err != nil {
		return "", "", errors.AddCode(err, "469884")
	}

	// add watermark
	if watermarkText != "" {
		srcImage, err = func(srcImage image.Image) (image.Image, error) {
			// Create a new image with the same size as the original image
			b := srcImage.Bounds()
			watermark := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))

			// Draw the original image onto the new image
			draw.Draw(watermark, watermark.Bounds(), srcImage, b.Min, draw.Src)

			// Load the font with a larger size
			face, err := func(size float64) (font.Face, error) {
				fontBytes, err := os.ReadFile(fmt.Sprintf("assets/fonts/%s", watermarkFont))
				if err != nil {
					return nil, errors.AddCode(err, "541757")
				}
				f, err := opentype.Parse(fontBytes)
				if err != nil {
					return nil, errors.AddCode(err, "804546")
				}
				return opentype.NewFace(f, &opentype.FaceOptions{
					Size:    size,
					DPI:     72,
					Hinting: font.HintingFull,
				})
			}(70)
			if err != nil {
				return nil, err
			}

			// calc width of the text
			width := func(face font.Face, text string) fixed.Int26_6 {
				var width fixed.Int26_6
				for _, r := range text {
					adv, ok := face.GlyphAdvance(r)
					if ok {
						width += adv
					}
				}
				return width
			}

			// Create a new image with the watermark text
			col := color.RGBA{0, 0, 0, 128} // Black color with transparency
			point := fixed.Point26_6{
				X: fixed.I(b.Dx())/2 - width(face, watermarkText)/2,
				Y: fixed.I(b.Dy() - 30),
			}

			d := &font.Drawer{
				Dst:  watermark,
				Src:  image.NewUniform(col),
				Face: face,
				Dot:  point,
			}
			d.DrawString(watermarkText)

			return watermark, nil
		}(srcImage)
	}

	// Resize the image
	resizedImage := resize.Resize(size, size, srcImage, resize.Lanczos3)

	// Create the output file for the resized JPG
	jpgFileName := filename + ".jpg"
	outFile, err := os.Create(filePath + jpgFileName)
	if err != nil {
		return "", "", errors.AddCode(err, "375715")
	}
	defer outFile.Close()

	// Encode the resized image to JPG
	err = jpeg.Encode(outFile, resizedImage, &jpeg.Options{Quality: int(compression)})
	if err != nil {
		return "", "", errors.AddCode(err, "334924")
	}

	// Create the output file for the WebP image
	webpFileName := filename + ".webp"
	webpFile, err := os.Create(filePath + webpFileName)
	if err != nil {
		return "", "", errors.AddCode(err, "849211")
	}
	defer webpFile.Close()

	// Encode the resized image to WebP
	err = webp.Encode(webpFile, resizedImage, &webp.Options{Lossless: false, Quality: float32(compression)})
	if err != nil {
		return "", "", errors.AddCode(err, "89555")
	}

	return jpgFileName, webpFileName, nil
}

func (s *Service) Get(ctx context.Context, filter *Filter) (*Item, error) {
	return s.repository.Get(ctx, filter)
}

func (s *Service) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	return s.repository.List(ctx, filter)
}

func (s *Service) Create(ctx context.Context, item *Item) (*uuid.UUID, error) {
	return s.repository.Create(ctx, item)
}

func (s *Service) Update(ctx context.Context, item *Item) error {
	return s.repository.Update(ctx, item)
}

func (s *Service) Patch(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) error {
	return s.repository.Patch(ctx, id, fields)
}

func (s *Service) UpdatedAt(ctx context.Context, id *uuid.UUID) (*time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *Service) TableIndexCount(ctx context.Context) (*uint64, error) {
	return s.repository.TableIndexCount(ctx)
}

func (s *Service) MaxSortOrder(ctx context.Context) (*uint64, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *Service) Delete(ctx context.Context, id *uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
