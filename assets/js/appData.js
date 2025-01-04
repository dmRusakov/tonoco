const a = this // a == applications

// functions
a.func = {}

// cache
a.cache = {}
a.cache.grid = {}

// browser
a.browser = (() => {
    return {
        userAgent: navigator.userAgent,
        platform: navigator.platform,
        isSupportWebp: (() => {
            const elem = document.createElement('canvas');
            if (!!(elem.getContext && elem.getContext('2d'))) {
                return elem.toDataURL('image/webp').indexOf('data:image/webp') == 0;
            }
            return false;
        })(),
    }
})()

// make top header
a.makeTopHeader = async () => {
    const dom = a.dom.querySelector(".topHeader")
    if (!dom) return false

    dom.pageTitile = dom.querySelector(".title")
    dom.menu = dom.querySelector(".menu")

    // func
    dom.setTitle = async (title) => {
        dom.pageTitile.innerHTML = title
    }

    // done
    return true
}

// make header
a.makeHeader = async () => {
    const dom = a.dom.querySelector(".header")
    if (!dom) return false

    dom.mobileMenu = dom.querySelector(".mobileMenu")
    dom.menu = dom.querySelector("nav.menu ul")
    dom.iconSize = dom.menu.querySelector("li a svg").getBoundingClientRect().width

    dom.menu.querySelectorAll("li a").forEach((a) => {
        a.icon = a.querySelector("svg")
        dom.menu[a.id] = a

        // mouse in event
        a.addEventListener("mouseover", () => {
            a.icon.style.width = dom.iconSize + 6 + "px"
            a.icon.style.margin = "-3px"
            a.icon.style.fill = "var(--menuIconHoverColor)"
        })

        // mouse out event
        a.addEventListener("mouseout", () => {
            a.icon.style.width = dom.iconSize + "px"
            a.icon.style.margin = "auto"
            a.icon.style.fill = "var(--menuIconColor)"
        })
    })

    // done
    return true
}

// make shop page
a.makeShopPage = async () => {
    const dom = a.dom.querySelector(".shop")
    if (!dom) return false

    await Promise.all([

        // title
        Promise.resolve().then(() => {
            dom.title = dom.querySelector(".header h1");
        }),

        // short description
        Promise.resolve().then(() => {
            dom.shortDescription = dom.querySelector(".header p");
        }),

        // description
        Promise.resolve().then(() => {
            dom.description = dom.querySelector("#description");
        }),

        // filters
        Promise.resolve().then(() => {
            dom.filters = dom.querySelector(".filters")
            dom.filters.style.removeProperty("display")
        }),

        // products grid
        Promise.resolve().then(() => {
            dom.products = dom.querySelector("#products .grid")
            dom.products.querySelectorAll("a").forEach(async (itemDom, i) => {
                const item = {
                    id: itemDom.getAttribute("id") || null,
                    brand: itemDom.querySelectorAll("h2 span")[0]?.innerHTML || null,
                    name: itemDom.querySelectorAll("h2 span")[1]?.innerHTML || null,
                    shortDescription: itemDom.querySelector("p.shortDescription").innerHTML,
                    status: itemDom.querySelector("div.dataContainer")?.getAttribute("status") || null,
                    salePrice: itemDom.querySelector("div.dataContainer span.salePrice")?.innerHTML || null,
                    regularPrice: itemDom.querySelector("div.dataContainer span.regularPrice")?.innerHTML || null,
                    sku: itemDom.querySelector("div.skuContainer p.sku")?.innerHTML || null,
                    tty: itemDom.querySelector("div.dataContainer")?.getAttribute("tty") || null,
                }

                // images
                item.images = {}
                itemDom.querySelector("picture").querySelectorAll("source").forEach((img) => {
                    item.images[img.getAttribute("source")] = img.getAttribute("srcset")
                })

                // save to cache
                a.cache.grid[item.id] = item

                // item count
                itemDom.querySelector(".skuContainer .itemCounter").innerHTML = "Item # " + (i + 1)

                const statusDom = itemDom.querySelector("div.dataContainer p.info")

                // mouse in event
                itemDom.addEventListener("mouseover", () => {
                    // picture
                    const picture = itemDom.querySelector("picture");
                    if (picture.classList.contains("hover")) return;
                    picture.classList.replace("main", "hover");

                    picture.querySelectorAll("source").forEach((imgSource) => {
                        const source = imgSource.getAttribute("source");
                        const [attr, newAttr] = ["main-webp", "main"].includes(source) ? ["srcset", "main-srcset"] : ["hover-srcset", "srcset"];
                        const src = imgSource.getAttribute(attr);
                        if (src) {
                            imgSource.setAttribute(newAttr, src);
                            imgSource.removeAttribute(attr);
                        }
                    });

                    // status
                    statusDom.innerHTML = item.status === "in_stock"
                        ? "Order now"
                        : item.status === "pre_order"
                            ? "Pre-order"
                            : "";
                });

                // mouse out event
                itemDom.addEventListener("mouseout", () => {
                    // picture
                    const picture = itemDom.querySelector("picture");
                    if (picture.classList.contains("main")) return;
                    picture.classList.replace("hover", "main");

                    picture.querySelectorAll("source").forEach((imgSource) => {
                        const source = imgSource.getAttribute("source");
                        const [attr, newAttr] = ["main-webp", "main"].includes(source) ? ["main-srcset", "srcset"] : ["srcset", "hover-srcset"];
                        const src = imgSource.getAttribute(attr);
                        if (src) {
                            imgSource.setAttribute(newAttr, src);
                            imgSource.removeAttribute(attr);
                        }
                    });

                    // status
                    statusDom.innerHTML = ""
                });

            })
        }),

        // products grid footer (pagination, count on page)
        Promise.resolve().then(() => {
            dom.productsFooter = dom.querySelector("#products .gridFooter")
        }),


        Promise.resolve().then(() => {}),
    ]);

    // dom










    // done
    return true
}