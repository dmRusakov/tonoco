document.addEventListener('DOMContentLoaded', async () => {
    // dom
    a.dom = document

    await Promise.all([
        a.makeTopHeader(),
        a.makeHeader(),
        a.makeShopPage()
    ]);

}, false);