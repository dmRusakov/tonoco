document.addEventListener('DOMContentLoaded', async () => {
    this.hd = {} //hd == header

    /** Dom **/
    this.hd.dom = {}
    this.hd.dom.header = document.querySelector(".mmHeader")
    this.hd.dom.topHeader = document.querySelector(".mmTopHeader")
    this.hd.dom.mobileMenu = this.hd.dom.header.querySelector(".mmMobileMenu")
    this.hd.dom.menu = this.hd.dom.header.querySelector("nav.menu")
    this.hd.dom.dashboardTab = this.hd.dom.menu.querySelector("a#mmDashboard")
    this.hd.dom.orderTab = this.hd.dom.menu.querySelector("a#mmOrders")
    this.hd.dom.productTab = this.hd.dom.menu.querySelector("a#mmProducts")
    this.hd.dom.categoryTab = this.hd.dom.menu.querySelector("a#mmCategories")
    this.hd.dom.pagesTab = this.hd.dom.menu.querySelector("a#mmPages")
    this.hd.dom.integrationTab = this.hd.dom.menu.querySelector("a#mmIntegration")
    this.hd.dom.cuponTab = this.hd.dom.menu.querySelector("a#mmCoupon")
    this.hd.dom.mediaTab = this.hd.dom.menu.querySelector("a#mmMedia")
    this.hd.dom.settingsTab = this.hd.dom.menu.querySelector("a#mmSettings")

    /** Functions **/
    this.hd.func = {}

    // make header element active
    this.hd.func.activePage = async () => {
        switch (this.param.url.pathname) {
            case '/':
                this.hd.dom.dashboardTab.classList.add("active")
                break
            case '/orders':
                this.hd.dom.orderTab.classList.add("active")
                break
            case '/grid':
                this.hd.dom.productTab.classList.add("active")
                break
            case '/categories':
                this.hd.dom.categoryTab.classList.add("active")
                break
            case '/pages':
                ad.hd.dom.pagesTab.classList.add("active")
                break
            case '/integration':
                this.hd.dom.integrationTab.classList.add("active")
                break
            case '/coupon':
                this.hd.dom.cuponTab.classList.add("active")
                break
            case '/media':
                this.hd.dom.mediaTab.classList.add("active")
                break
            case '/settings':
                this.hd.dom.settingsTab.classList.add("active")
                break
        }
    }

    // toggle mobile menu
    this.hd.func.toggleMobileMenu = async (action = null) => {
        switch (action) {
            case "show":
                this.hd.dom.menu.style.display = "block"
                this.hd.dom.mobileMenu.classList.add("active")
                break
            case "hide":
                this.hd.dom.menu.style.display = "none"
                this.hd.dom.mobileMenu.classList.remove("active")
                break
            default:
                if (this.hd.dom.menu.style.display === "none") {
                    this.hd.func.toggleMobileMenu("show")
                } else {
                    this.hd.func.toggleMobileMenu("hide")
                }
        }
    }

    /** Actions **/
    // add active class to header element
    this.hd.func.activePage()
    this.hd.dom.mobileMenu.addEventListener("click", this.hd.func.toggleMobileMenu, false)

    // update Item count in cart(gridItems)


}, false);