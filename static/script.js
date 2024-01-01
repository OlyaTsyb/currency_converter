function showExitButtons() {
    var exitButtonsContainer = document.getElementById('exitButtons');

    if (exitButtonsContainer) {
        exitButtonsContainer.style.display = 'block';

        var yesBtn = document.getElementById('yesBtn');
        var noBtn = document.getElementById('noBtn');

        if (yesBtn) yesBtn.disabled = false;
        if (noBtn) noBtn.disabled = false;

        if (yesBtn) {
            yesBtn.addEventListener('click', function () {
                if (exitButtonsContainer) {
                    exitButtonsContainer.style.display = 'none';
                    window.location.href = "/index";
                }
            });
        }

        if (noBtn) {
            noBtn.addEventListener('click', function () {
                window.location.href = "/";
            });
        }
    }
}

document.getElementById('convertButton').addEventListener('click', function () {
    var fromCurrency = document.getElementById('fromCurrency').value;
    var toCurrency = document.getElementById('toCurrency').value;
    var date = document.getElementById('date').value;
    var amount = document.getElementById('amount').value;

    fetch(`/convert?from=${fromCurrency}&to=${toCurrency}&date=${date}&amount=${amount}`)
        .then(response => response.text())
        .then(result => {
            document.getElementById('result').innerHTML = result;

            if (!result.includes("Invalid")) {
                showExitButtons();
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
});
