document.addEventListener('DOMContentLoaded', async () => {



    // product grid
    document.querySelectorAll(".gridItems").forEach((gridItems) => {
        let count = 0;
        gridItems.querySelectorAll(".itemNo").forEach((itemNo) => {
            itemNo.textContent = "Item #" + ++count
        })
    })
})