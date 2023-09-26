document.addEventListener('DOMContentLoaded', async () => {
    appData.header = {}

    /** Dom **/
    appData.header.dom = {}
    appData.header.dom.header = document.querySelector(".mmHeader")
    appData.header.dom.topHeader = document.querySelector(".mmTopHeader")
    appData.header.dom.menu = appData.header.dom.header.querySelector("nav.menu")
    appData.header.dom.dashboardTab = appData.header.dom.menu.querySelector("a#mmDashboard")
    appData.header.dom.orderTab = appData.header.dom.menu.querySelector("a#mmOrders")
    appData.header.dom.productTab = appData.header.dom.menu.querySelector("a#mmProducts")
    appData.header.dom.categoryTab = appData.header.dom.menu.querySelector("a#mmCategories")
    appData.header.dom.pagesTab = appData.header.dom.menu.querySelector("a#mmPages")
    appData.header.dom.integrationTab = appData.header.dom.menu.querySelector("a#mmIntegration")
    appData.header.dom.cuponTab = appData.header.dom.menu.querySelector("a#mmCoupon")
    appData.header.dom.mediaTab = appData.header.dom.menu.querySelector("a#mmMedia")
    appData.header.dom.settingsTab = appData.header.dom.menu.querySelector("a#mmSettings")

    /** Functions **/
    appData.header.func = {}

    // make header element active
    appData.header.func.activePage = async () => {
        switch (appData.param.url.pathname) {
            case '/':
                appData.header.dom.dashboardTab.classList.add("active")
                break;
            case '/orders':
                appData.header.dom.orderTab.classList.add("active")
                break;
            case '/products':
                appData.header.dom.productTab.classList.add("active")
                break;
            case '/categories':
                appData.header.dom.categoryTab.classList.add("active")
                break;
            case '/pages':
                appData.header.dom.pagesTab.classList.add("active")
                break;
            case '/integration':
                appData.header.dom.integrationTab.classList.add("active")
                break;
            case '/coupon':
                appData.header.dom.cuponTab.classList.add("active")
                break;
            case '/media':
                appData.header.dom.mediaTab.classList.add("active")
                break;
            case '/settings':
                appData.header.dom.settingsTab.classList.add("active")
                break;
        }
    }

    /** Actions **/
    // add active class to header element
    appData.header.func.activePage()

}, false);