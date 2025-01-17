class TopHeader {
    constructor() {
        self.topHeader = this
    }
    
    async init() {
        this.dom = document.querySelector('.topHeader')
        this.pageTitile = this.dom.querySelector('.title')
        this.menu = this.dom.querySelector('.menu')
    }

    setTitle(title) {
        this.pageTitile.innerHTML = title
    }
}

class Header {
    constructor() {
        self.header = this
    }

    async init() {
        this.dom = document.querySelector('.header')
        this.mobileMenu = this.dom.querySelector('.mobileMenu')
        this.menu = this.dom.querySelector('nav.menu ul')
        this.iconSize = this.menu.querySelector('li a svg').getBoundingClientRect().width

        this.menu.querySelectorAll('li a').forEach((a) => {
            a.icon = a.querySelector('svg')
            this.menu[a.id] = a

            a.addEventListener('mouseover', () => {
                a.icon.style.width = this.iconSize + 6 + 'px'
                a.icon.style.margin = '-3px'
                a.icon.style.fill = 'var(--menuIconHoverColor)'
            })

            a.addEventListener('mouseout', () => {
                a.icon.style.width = this.iconSize + 'px'
                a.icon.style.margin = 'auto'
                a.icon.style.fill = 'var(--menuIconColor)'
            })
        })
    }
}

new TopHeader()
new Header()

document.addEventListener('DOMContentLoaded', async () => {
    await self.topHeader.init()
    await self.header.init()
}, false);