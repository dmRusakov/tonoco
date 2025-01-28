// class Shop
class Shop {
    constructor() {
        self.shop = this
    }

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
            new Promise(async (resolve) => {
                const dom = this.filters = this.header.querySelector(".filters");
                dom.style.removeProperty("display");
                let isFileterExist = false;

                this.filters.querySelectorAll(".filter").forEach((filterDom) => {
                    const filter = {
                        id: filterDom.querySelector("select").getAttribute("name"),
                        title: filterDom.querySelector("label").innerHTML,
                        options: {}
                    }

                    filterDom.querySelector("select").querySelectorAll("option").forEach((optionDom) => {
                        const option = {
                            id: optionDom.getAttribute("value-id"),
                            title: optionDom.innerHTML,
                            key: optionDom.getAttribute("value"),
                        }

                        filter.options[option.id] = option;
                    })

                    // save to cache
                    self.app.setFilterCache(filter.id, filter)
                    isFileterExist = true;
                })

                // Choices
                if (isFileterExist) {
                    // add Choices CDN to header
                    const head = document.head;
                    const link = document.createElement("link");
                    link.rel = "stylesheet";
                    link.href = "https://cdn.jsdelivr.net/npm/choices.js/public/assets/styles/choices.min.css";
                    head.appendChild(link);

                    const script = document.createElement("script");
                    script.src = "https://cdn.jsdelivr.net/npm/choices.js/public/assets/scripts/choices.min.js";
                    head.appendChild(script);

                    // wait for Choices
                    await new Promise((resolve) => {
                        script.onload = resolve;
                    });

                    this.filters.querySelectorAll(".filter select").forEach((filterDom) => {
                        new Choices(filterDom, {
                            removeItemButton: true,
                            placeholder: true,
                        });
                    })
                }

                resolve();
            }),

            // grid
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
                    self.app.setGridCache(item.id, item)

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

            // perPage
            new Promise((resolve) => {
                const dom = this.perPage = this.footer.querySelector("select");
                dom.set = async (t) => {
                    dom.value = t;
                }
                dom.get = async () => {
                    return dom.value;
                }
                resolve();
            }),

            // onPage
            new Promise((resolve) => {
                const dom = this.onPage = this.footer.querySelector("span");
                dom.set = async (t) => {
                    dom.value = t;
                }
                resolve();
            }),

            // pagination
            new Promise((resolve) => {
                const dom = this.pagination = this.footer.querySelector(".pagination");
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

new Shop()
document.addEventListener('DOMContentLoaded', async () => {
    await self.shop.init()
})