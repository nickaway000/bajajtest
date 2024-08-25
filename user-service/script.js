document.getElementById('bfhlForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    const dataInput = document.getElementById('data').value;
    const dataArray = dataInput.split(',').map(item => item.trim());

    const jsonData = {
        data: dataArray
    };

    try {
        const response = await fetch('http://localhost:8080/bfhl', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonData)
        });

        const responseData = await response.json();

        document.getElementById('response').innerText = JSON.stringify(responseData, null, 4);
    } catch (error) {
        document.getElementById('response').innerText = 'Error: ' + error.message;
    }
});
