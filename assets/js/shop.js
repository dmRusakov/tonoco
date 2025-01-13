// class Shop
class Shop {
    async init() {
        this.dom = document.querySelector('.shop')
        this.header = this.dom.querySelector('.header')
        this.footer = this.dom.querySelector('.gridFooter')

        await Promise.all([

            // name
            new Promise((resolve) => {
                const dom = this.name = this.header.querySelector("#title h1");
                dom.set = async (t) => {
                    this.name.innerText = t;
                }
                resolve();
            }),

            // shortDescription
            new Promise((resolve) => {
                const dom = this.shortDescription = this.header.querySelector("#title p");
                dom.set = async (t) => {
                    this.name.innerText = t;
                }
                resolve();
            }),

            // description
            new Promise((resolve) => {
                const dom = this.description = this.dom.querySelector("#description");
                dom.set = async (t) => {
                    this.name.innerText = t;
                }
                resolve();
            }),

            // filters
            new Promise((resolve) => {
                const dom = this.filters = this.header.querySelector(".filters");
                dom.style.removeProperty("display");
                resolve();
            }),

            // products
            new Promise((resolve) => {
                const dom = this.products = this.dom.querySelector("#products .grid");
                dom.querySelectorAll("a").forEach(async (itemDom, i) => {
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
                resolve();
            }),

            // count items
            new Promise((resolve) => {
                const dom = this.footer = this.header.querySelector(".countItems");
                dom.perPage = this.footer.querySelector("select");
                dom.perPage.set = async (t) => {
                    dom.perPage.value = t;
                }
                dom.perPage.get = async () => {
                    return dom.perPage.value;
                }
                dom.onPage = this.footer.querySelector("span");
                dom.onPage.set = async (t) => {
                    dom.onPage.innerHTML = t;
                }
                resolve();
            }),

            // pagination
            new Promise((resolve) => {
                const dom = this.footer = this.header.querySelector(".pagination");
                dom.pages = dom.querySelectorAll("a");
                dom.pages.forEach((page) => {
                    page.set = async (t) => {
                        page.innerHTML = t;
                    }
                })
                resolve();
            }),
        ])
    }
}

const shop = self.shop = new Shop()
document.addEventListener('DOMContentLoaded', async () => {
    await shop.init()
    console.log(shop)
})