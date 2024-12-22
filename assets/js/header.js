document.addEventListener('DOMContentLoaded', async () => {
    // dom
    a.dom = document

    // top header
    a.topHeader = a.dom.querySelector(".topHeader")
    const topHeaderReady = a.func.makeTopHeader(a.topHeader)

    // header
    a.header = a.dom.querySelector(".header")
    const headerReady = a.func.makeHeader(a.header)








    // wait all ready
     await Promise.all([topHeaderReady, headerReady])

    console.log(a.header.menu["mmProducts"])
}, false);