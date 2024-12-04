package shop_page

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/google/uuid"
	"sync"
)

func (u *UseCase) GetShopPage(ctx context.Context, pageUrl string) (*pages.ShopPage, []error) {
	var wg sync.WaitGroup
	var errs []error

	// ger shop page by url
	pageIds, err := u.shopPage.Get(ctx, &db.ShopPageFilter{
		Urls: &[]string{pageUrl},
	})

	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	page := pages.ShopPage{
		Id:        pageIds.Id,
		Url:       pageIds.Url,
		Page:      pageIds.Page,
		PerPage:   pageIds.PerPage,
		SortOrder: pageIds.SortOrder,
		Active:    pageIds.Active,
		Prime:     pageIds.Prime,
		CreatedAt: pageIds.CreatedAt,
		CreatedBy: pageIds.CreatedBy,
		UpdatedAt: pageIds.UpdatedAt,
		UpdatedBy: pageIds.UpdatedBy,
	}

	// Name
	wg.Add(1)
	go func() {
		defer wg.Done()

		// check for nil
		if pageIds.Name == uuid.Nil {
			page.Name = ""
			return
		}

		// get text
		text, e := u.text.Get(ctx, &db.TextFilter{
			Source:    pointer.StringToPtr("shop_page"),
			SubSource: pointer.StringToPtr("name"),
			Language:  pointer.StringToPtr(u.cfg.AppDefaultLanguage),
			Ids:       &[]uuid.UUID{pageIds.Name},
		})

		if e != nil {
			errs = append(errs, e)
			page.Name = ""
			return
		}

		page.Name = text.Text
	}()

	// SeoTitle
	wg.Add(1)
	go func() {
		defer wg.Done()

		// check for nil
		if pageIds.SeoTitle == uuid.Nil {
			page.SeoTitle = ""
			return
		}

		// get text
		text, e := u.text.Get(ctx, &db.TextFilter{
			Source:    pointer.StringToPtr("shop_page"),
			SubSource: pointer.StringToPtr("seo_title"),
			Language:  pointer.StringToPtr(u.cfg.AppDefaultLanguage),
			Ids:       &[]uuid.UUID{pageIds.SeoTitle},
		})

		if e != nil {
			errs = append(errs, e)
			page.SeoTitle = ""
			return
		}

		page.SeoTitle = text.Text
	}()

	// ShortDescription
	wg.Add(1)
	go func() {
		defer wg.Done()

		// check for nil
		if pageIds.ShortDescription == uuid.Nil {
			page.ShortDescription = ""
			return
		}

		text, e := u.text.Get(ctx, &db.TextFilter{
			Source:    pointer.StringToPtr("shop_page"),
			SubSource: pointer.StringToPtr("short_description"),
			Language:  pointer.StringToPtr(u.cfg.AppDefaultLanguage),
			Ids:       &[]uuid.UUID{pageIds.ShortDescription},
		})

		if e != nil {
			errs = append(errs, e)
			page.ShortDescription = ""
			return
		}

		page.ShortDescription = text.Text
	}()

	// Description
	wg.Add(1)
	go func() {
		defer wg.Done()

		// check for nil
		if pageIds.Description == uuid.Nil {
			page.Description = ""
			return
		}

		text, e := u.text.Get(ctx, &db.TextFilter{
			Source:    pointer.StringToPtr("shop_page"),
			SubSource: pointer.StringToPtr("description"),
			Language:  pointer.StringToPtr(u.cfg.AppDefaultLanguage),
			Ids:       &[]uuid.UUID{pageIds.Description},
		})

		if e != nil {
			errs = append(errs, e)
			page.Description = ""
			return
		}

		page.Description = text.Text
	}()

	// Image
	wg.Add(1)
	go func() {
		defer wg.Done()

		// check for nil
		if pageIds.ImageId == uuid.Nil {
			page.Image = nil
			return
		}

		image, e := u.image.Get(ctx, &db.ImageFilter{
			Ids: &[]uuid.UUID{pageIds.ImageId},
		})

		page.Image = image

		if e != nil {
			errs = append(errs, e)
		}
	}()

	// HoverImage
	wg.Add(1)
	go func() {
		defer wg.Done()

		// check for nil
		if pageIds.HoverImageId == uuid.Nil {
			page.HoverImage = nil
			return
		}

		image, e := u.image.Get(ctx, &db.ImageFilter{
			Ids: &[]uuid.UUID{pageIds.HoverImageId},
		})

		page.HoverImage = image

		if e != nil {
			errs = append(errs, e)
		}
	}()

	// wait for all goroutines
	wg.Wait()

	return &page, nil
}
