document.addEventListener('DOMContentLoaded', async () => {
    document.querySelectorAll(".gridContainer").forEach((gridContainer) => {
        let count = 0;
        gridContainer.querySelectorAll(".itemNo").forEach((itemNo) => {
            itemNo.textContent = "Item #" + ++count
        })
    })
})