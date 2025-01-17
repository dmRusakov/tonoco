const a = this // a == applications

// cache
self.cache = {}
self.cache.grid = {}
self.cache.filter = {}

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

        // products grid
        Promise.resolve().then(() => {
            dom.products = dom.querySelector("#products .grid")

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