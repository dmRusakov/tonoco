document.addEventListener('DOMContentLoaded', async () => {
    ad.hd = {} //hd == header

    /** Dom **/
    ad.hd.dom = {}
    ad.hd.dom.header = document.querySelector(".mmHeader")
    ad.hd.dom.topHeader = document.querySelector(".mmTopHeader")
    ad.hd.dom.mobileMenu = ad.hd.dom.header.querySelector(".mmMobileMenu")
    ad.hd.dom.menu = ad.hd.dom.header.querySelector("nav.menu")
    ad.hd.dom.dashboardTab = ad.hd.dom.menu.querySelector("a#mmDashboard")
    ad.hd.dom.orderTab = ad.hd.dom.menu.querySelector("a#mmOrders")
    ad.hd.dom.productTab = ad.hd.dom.menu.querySelector("a#mmProducts")
    ad.hd.dom.categoryTab = ad.hd.dom.menu.querySelector("a#mmCategories")
    ad.hd.dom.pagesTab = ad.hd.dom.menu.querySelector("a#mmPages")
    ad.hd.dom.integrationTab = ad.hd.dom.menu.querySelector("a#mmIntegration")
    ad.hd.dom.cuponTab = ad.hd.dom.menu.querySelector("a#mmCoupon")
    ad.hd.dom.mediaTab = ad.hd.dom.menu.querySelector("a#mmMedia")
    ad.hd.dom.settingsTab = ad.hd.dom.menu.querySelector("a#mmSettings")

    /** Functions **/
    ad.hd.func = {}

    // make header element active
    ad.hd.func.activePage = async () => {
        switch (ad.param.url.pathname) {
            case '/':
                ad.hd.dom.dashboardTab.classList.add("active")
                break
            case '/orders':
                ad.hd.dom.orderTab.classList.add("active")
                break
            case '/products':
                ad.hd.dom.productTab.classList.add("active")
                break
            case '/categories':
                ad.hd.dom.categoryTab.classList.add("active")
                break
            case '/pages':
                ad.hd.dom.pagesTab.classList.add("active")
                break
            case '/integration':
                ad.hd.dom.integrationTab.classList.add("active")
                break
            case '/coupon':
                ad.hd.dom.cuponTab.classList.add("active")
                break
            case '/media':
                ad.hd.dom.mediaTab.classList.add("active")
                break
            case '/settings':
                ad.hd.dom.settingsTab.classList.add("active")
                break
        }
    }

    // toggle mobile menu
    ad.hd.func.toggleMobileMenu = async (action = null) => {
        switch (action) {
            case "show":
                ad.hd.dom.menu.style.display = "block"
                ad.hd.dom.mobileMenu.classList.add("active")
                break
            case "hide":
                ad.hd.dom.menu.style.display = "none"
                ad.hd.dom.mobileMenu.classList.remove("active")
                break
            default:
                if (ad.hd.dom.menu.style.display == "none") {
                    ad.hd.func.toggleMobileMenu("show")
                } else {
                    ad.hd.func.toggleMobileMenu("hide")
                }
        }
    }

    /** Actions **/
    // add active class to header element
    ad.hd.func.activePage()
    ad.hd.dom.mobileMenu.addEventListener("click", ad.hd.func.toggleMobileMenu, false)

}, false);