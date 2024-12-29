document.addEventListener('DOMContentLoaded', async () => {
    // dom
    a.dom = document

    await Promise.all([
        a.func.makeTopHeader(),
        a.func.makeHeader(),
    ]);

}, false);