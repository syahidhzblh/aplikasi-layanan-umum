document.getElementById('btnQuery').addEventListener('click', () => {
    fetch('/run-query')
        .then(response => response.json())
        .then(data => {
            console.log('Query Result:', data);
            // You can display the result on the page if needed
        })
        .catch(error => console.error('Error:', error));
});