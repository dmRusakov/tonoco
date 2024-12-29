const a = this // a == applications

// functions
a.func = {}

// make top header
a.func.makeTopHeader = async () => {
    const dom = a.dom.querySelector(".topHeader")
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
a.func.makeHeader = async () => {
    const dom = a.dom.querySelector(".header")

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