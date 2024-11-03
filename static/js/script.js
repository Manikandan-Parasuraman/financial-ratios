function showForm(formId) {
    document.querySelectorAll('.form-container').forEach(form => form.style.display = 'none');
    document.getElementById(formId).style.display = "block";
}

function calculateEmergencyFund() {
    const emergencyFund = document.getElementById("emergency_fund").value;
    const monthlyExpenses = document.getElementById("monthly_expenses").value;

    fetch("/calculate-emergency-fund", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `emergency_fund=${emergencyFund}&monthly_expenses=${monthlyExpenses}`
    })
        .then(response => response.json())
        .then(data => {
            const resultDiv = document.getElementById("result");
            resultDiv.innerHTML = `Emergency Fund Ratio: ${data.months} months`;
            resultDiv.className = data.recommendation;
        });
}

function calculateSavingsToIncome() {
    const savings = document.getElementById("savings").value;
    const grossIncome = document.getElementById("gross_income").value;

    fetch("/calculate-savings-to-income", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `savings=${savings}&gross_income=${grossIncome}`
    })
        .then(response => response.json())
        .then(data => {
            const resultDiv = document.getElementById("resultSavings");
            resultDiv.innerHTML = `Savings to Income Ratio: ${data.percentage.toFixed(2)}%`;
            resultDiv.className = data.recommendation;
        });
}

function calculateDebtToIncome() {
    const totalDebt = document.getElementById("total_debt").value;
    const grossIncome = document.getElementById("gross_income").value;

    fetch("/calculate-debt-to-income", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `total_debt=${totalDebt}&gross_income=${grossIncome}`
    })
        .then(response => response.json())
        .then(data => {
            const resultDiv = document.getElementById("resultSavings");
            resultDiv.innerHTML = `Savings to Income Ratio: ${data.percentage.toFixed(2)}%`;
            resultDiv.className = data.recommendation;
        });
}