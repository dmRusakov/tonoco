document.addEventListener('DOMContentLoaded', async () => {
    document.querySelectorAll(".gridItems").forEach((gridItems) => {
        let count = 0;
        gridItems.querySelectorAll(".itemNo").forEach((itemNo) => {
            itemNo.textContent = "Item #" + ++count
        })
    })
})