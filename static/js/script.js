// Function to navigate to specific calculator and show its form
function navigateToCalculator(calculatorType) {
    // First navigate to calculator page
    window.location.href = '/';
    
    // Store the calculator type in session storage
    sessionStorage.setItem('selectedCalculator', calculatorType);
}

// Function to show selected form and hide others
function showForm(formId) {
    // Hide all forms with fade out
    const forms = document.querySelectorAll('.form-container');
    forms.forEach(form => {
        form.style.display = 'none';
        form.classList.remove('show');
    });

    // Show selected form with fade in
    const selectedForm = document.getElementById(formId);
    if (selectedForm) {
        selectedForm.style.display = 'block';
        // Trigger reflow
        selectedForm.offsetHeight;
        selectedForm.classList.add('show');
    }

    // Clear previous results with animation
    const resultDiv = document.getElementById('result');
    if (resultDiv) {
        resultDiv.style.opacity = '0';
        setTimeout(() => {
            resultDiv.innerHTML = '';
            resultDiv.style.opacity = '1';
        }, 300);
    }

    // Update active button state with smooth transition
    const buttons = document.querySelectorAll('.sidebar-button');
    buttons.forEach(button => {
        button.classList.remove('active');
    });

    const activeButton = document.querySelector(`[onclick="showForm('${formId}')"]`);
    if (activeButton) {
        activeButton.classList.add('active');
    }

    // Store the calculator type in session storage
    sessionStorage.setItem('selectedCalculator', formId);
}

// Show emergency fund form by default and set up initial state
document.addEventListener('DOMContentLoaded', function() {
    // Show emergency fund form by default
    const defaultForm = document.getElementById('emergencyFundForm');
    if (defaultForm) {
        defaultForm.style.display = 'block';
        defaultForm.classList.add('show');
    }

    // Set the emergency fund button as active
    const defaultButton = document.querySelector('[onclick="showForm(\'emergencyFundForm\')"]');
    if (defaultButton) {
        defaultButton.classList.add('active');
    }
});

// Function to display calculation result with animation
function displayResult(result, type = 'success') {
    const resultDiv = document.getElementById('result');
    let interpretationText = '';
    let ratioValue = '';

    // Format the result based on the calculator type
    if (typeof result === 'object') {
        ratioValue = result.ratio;
        interpretationText = result.interpretation;
    } else {
        ratioValue = result;
    }

    // Create result display with animation
    const resultHTML = `
        <div class="result-display ${type}">
            <div class="ratio-value">${ratioValue}</div>
            ${interpretationText ? `<div class="interpretation">${interpretationText}</div>` : ''}
        </div>
    `;

    // Animate the result display
    resultDiv.style.opacity = '0';
    setTimeout(() => {
        resultDiv.innerHTML = resultHTML;
        // Trigger reflow
        resultDiv.offsetHeight;
        resultDiv.style.opacity = '1';
        
        // Add show class to animate the result display
        const displayElement = resultDiv.querySelector('.result-display');
        if (displayElement) {
            displayElement.classList.add('show');
        }
    }, 300);

    // Scroll the results into view
    resultDiv.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
}

// Function to clear result
function clearResult() {
    const resultDiv = document.getElementById('result');
    resultDiv.style.opacity = '0';
    setTimeout(() => {
        resultDiv.innerHTML = '';
        resultDiv.style.opacity = '1';
    }, 300);
    resultDiv.className = 'result-display';
}

// Update API endpoints to include /calculator prefix
const API_ENDPOINTS = {
    'emergency-fund': '/calculator/calculate-emergency-fund',
    'savings-ratio': '/calculator/calculate-savings-to-income',
    'debt-ratio': '/calculator/calculate-debt-to-income',
    'housing-ratio': '/calculator/calculate-housing-expenses',
    'net-worth-ratio': '/calculator/calculate-net-worth-income',
    'investment-ratio': '/calculator/calculate-investment-ratio',
    'retirement-ratio': '/calculator/calculate-retirement-savings',
    'liquidity-ratio': '/calculator/calculate-liquidity-fund'
};

function calculateEmergencyFund() {
    const emergencyFund = document.getElementById("emergency_fund").value;
    const monthlyExpenses = document.getElementById("monthly_expenses").value;

    if (!emergencyFund || !monthlyExpenses) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    fetch(API_ENDPOINTS['emergency-fund'], {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `emergency_fund=${emergencyFund}&monthly_expenses=${monthlyExpenses}`
    })
    .then(response => {
        if (!response.ok) throw new Error('Network response was not ok');
        return response.json();
    })
    .then(data => {
        if (data.error) {
            displayResult(`Error: ${data.error}`, "error");
        } else {
            displayResult(`Emergency Fund Ratio: ${data.months} months\n\n
                         <strong>Recommendation:</strong>\n
                         ${getRecommendation('emergency', data.months)}`, 
                         data.recommendation);
        }
    })
    .catch(error => {
        displayResult(`Error: ${error.message}`, "error");
    });
}

function calculateSavingsToIncome() {
    const savings = document.getElementById("savings").value;
    const grossIncome = document.getElementById("gross_income").value;

    if (!savings || !grossIncome) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    fetch(API_ENDPOINTS['savings-ratio'], {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `savings=${savings}&gross_income=${grossIncome}`
    })
    .then(response => {
        if (!response.ok) throw new Error('Network response was not ok');
        return response.json();
    })
    .then(data => {
        if (data.error) {
            displayResult(`Error: ${data.error}`, "error");
        } else {
            displayResult(`Savings to Income Ratio: ${data.percentage.toFixed(2)}%\n\n
                         <strong>Recommendation:</strong>\n
                         ${getRecommendation('savings', data.percentage)}`, 
                         data.recommendation);
        }
    })
    .catch(error => {
        displayResult(`Error: ${error.message}`, "error");
    });
}

function calculateDebtToIncome() {
    const totalDebt = document.getElementById("total_monthly_debt").value;
    const grossIncome = document.getElementById("gross_income_debt").value;

    console.log('Input values:', {
        totalDebt: totalDebt,
        grossIncome: grossIncome
    });

    if (!totalDebt || !grossIncome) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    // Create URLSearchParams for proper form encoding
    const formData = new URLSearchParams();
    formData.append('total_monthly_debt', totalDebt);
    formData.append('gross_income', grossIncome);

    console.log('Sending request with data:', formData.toString());

    fetch(API_ENDPOINTS['debt-ratio'], {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: formData.toString()
    })
    .then(response => {
        console.log('Response status:', response.status);
        return response.json().then(data => {
            console.log('Response data:', data);
            if (!response.ok) {
                throw new Error(data.error || 'Network response was not ok');
            }
            return data;
        });
    })
    .then(data => {
        console.log('Success response:', data);
        if (data.error) {
            displayResult(`Error: ${data.error}`, "error");
        } else {
            displayResult(`Debt to Income Ratio: ${data.percentage.toFixed(2)}%\n\n
                         <strong>Recommendation:</strong>\n
                         ${data.message}`, 
                         data.recommendation);
        }
    })
    .catch(error => {
        console.error('Error details:', error);
        displayResult(`Error: ${error.message}`, "error");
    });
}

function getRecommendation(type, value) {
    switch(type) {
        case 'emergency':
            if (value < 3) return "Your emergency fund is low. Aim to save at least 3-6 months of expenses.";
            if (value < 6) return "Good start! Consider building up to 6 months of expenses for better security.";
            return "Excellent! You have a strong emergency fund.";
        
        case 'savings':
            if (value < 10) return "Try to save at least 10% of your income for long-term financial health.";
            if (value < 20) return "Good saving habits! Consider increasing to 20% for faster wealth building.";
            return "Excellent saving rate! You're on track for strong financial growth.";
        
        case 'debt':
            if (value > 36) return "Your debt-to-income ratio is high. Consider debt reduction strategies.";
            if (value > 28) return "Watch your debt levels. Try to reduce non-essential borrowing.";
            return "Your debt-to-income ratio is in a healthy range.";
        
        case 'housing':
            if (value > 35) return "Your housing expenses are significantly high. Consider ways to reduce housing costs.";
            if (value > 28) return "Your housing expenses are slightly above recommended levels. Consider budgeting options.";
            return "Your housing expenses are within the recommended range (below 28%).";

        case 'networth':
            if (value < 0) return "Focus on debt reduction and building savings to improve your net worth.";
            if (value < 100) return "Work on building your net worth to at least match your annual income.";
            if (value < 300) return "Good progress! Your net worth exceeds your annual income.";
            return "Excellent financial position! Your net worth is significantly higher than your income.";
        
        case 'investment':
            if (value < 20) return "Consider increasing your investments to build long-term wealth.";
            if (value < 40) return "Good job! You have a healthy portion of your net worth invested.";
            if (value < 60) return "Excellent! You have a strong investment portfolio relative to your net worth.";
            return "While having substantial investments is good, consider maintaining some liquid assets for emergencies.";
        
        default:
            return "";
    }
}

function calculateNetWorthToIncome() {
    const netWorth = document.getElementById("net_worth").value;
    const annualIncome = document.getElementById("annual_income").value;

    console.log('Input values:', { netWorth, annualIncome });

    if (!netWorth || !annualIncome) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    const formData = new URLSearchParams();
    formData.append('net_worth', netWorth);
    formData.append('annual_income', annualIncome);

    console.log('Sending request with data:', formData.toString());

    fetch(API_ENDPOINTS['net-worth-ratio'], {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: formData.toString()
    })
    .then(response => {
        console.log('Response status:', response.status);
        return response.json().then(data => {
            console.log('Response data:', data);
            if (!response.ok) {
                throw new Error(data.error || 'Server error');
            }
            return data;
        });
    })
    .then(data => {
        console.log('Processing response:', data);
        if (data.status === "error") {
            throw new Error(data.error);
        }
        
        const resultHtml = `
            Net Worth to Income Ratio: ${data.percentage.toFixed(2)}%
            <br><br>
            <strong>Recommendation:</strong><br>
            ${data.message}
        `;
        
        displayResult(resultHtml, data.recommendation);
    })
    .catch(error => {
        console.error('Error:', error);
        displayResult(error.message || "An error occurred while calculating the ratio. Please try again.", "error");
    });
}

function calculateHousingExpenses() {
    const housingCost = document.getElementById("monthly_housing_cost").value;
    const grossIncome = document.getElementById("monthly_gross_income_housing").value;

    console.log('Input values:', {
        housingCost: housingCost,
        grossIncome: grossIncome
    });

    if (!housingCost || !grossIncome) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    const formData = new URLSearchParams();
    formData.append('monthly_housing_cost', housingCost);
    formData.append('monthly_gross_income', grossIncome);

    console.log('Sending request with data:', formData.toString());

    fetch(API_ENDPOINTS['housing-ratio'], {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: formData.toString()
    })
    .then(response => {
        console.log('Response status:', response.status);
        return response.json().then(data => {
            console.log('Response data:', data);
            if (!response.ok) {
                throw new Error(data.error || 'Network response was not ok');
            }
            return data;
        });
    })
    .then(data => {
        console.log('Success response:', data);
        if (data.error) {
            displayResult(`Error: ${data.error}`, "error");
        } else {
            displayResult(`Housing Expenses Ratio: ${data.percentage.toFixed(2)}%\n\n
                         <strong>Recommendation:</strong>\n
                         ${getRecommendation('housing', data.percentage)}`, 
                         data.recommendation);
        }
    })
    .catch(error => {
        console.error('Error details:', error);
        displayResult(`Error: ${error.message}`, "error");
    });
}

function calculateInvestmentRatio() {
    const totalInvestment = document.getElementById("total_investment").value;
    const netWorth = document.getElementById("net_worth_investment").value;

    console.log('Input values:', { totalInvestment, netWorth });

    if (!totalInvestment || !netWorth) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    const formData = new URLSearchParams();
    formData.append('total_investment', totalInvestment);
    formData.append('net_worth', netWorth);

    console.log('Sending request with data:', formData.toString());

    fetch(API_ENDPOINTS['investment-ratio'], {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: formData.toString()
    })
    .then(response => {
        console.log('Response status:', response.status);
        return response.json().then(data => {
            console.log('Response data:', data);
            if (!response.ok) {
                throw new Error(data.error || 'Server error');
            }
            return data;
        });
    })
    .then(data => {
        console.log('Processing response:', data);
        if (data.status === "error") {
            throw new Error(data.error);
        }
        
        const resultHtml = `
            Investment Ratio: ${data.percentage.toFixed(2)}%
            <br><br>
            <strong>Recommendation:</strong><br>
            ${data.message}
        `;
        
        displayResult(resultHtml, data.recommendation);
    })
    .catch(error => {
        console.error('Error:', error);
        displayResult(error.message || "An error occurred while calculating the ratio. Please try again.", "error");
    });
}

function calculateRetirementSavings() {
    const retirementSavings = document.getElementById('retirement_savings').value;
    const annualIncome = document.getElementById('annual_income_retirement').value;

    if (!retirementSavings || !annualIncome) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    const formData = new URLSearchParams();
    formData.append('retirement_savings', retirementSavings);
    formData.append('annual_income_retirement', annualIncome);

    fetch(API_ENDPOINTS['retirement-ratio'], {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: formData.toString()
    })
    .then(response => {
        return response.json().then(data => {
            if (!response.ok) {
                throw new Error(data.error || 'Network response was not ok');
            }
            return data;
        });
    })
    .then(data => {
        if (data.error) {
            displayResult(`Error: ${data.error}`, "error");
        } else {
            const resultHtml = `
                Retirement Savings Ratio: ${data.percentage.toFixed(2)}%
                <br><br>
                <strong>Recommendation:</strong><br>
                ${data.message}
            `;
            displayResult(resultHtml, data.recommendation);
        }
    })
    .catch(error => {
        displayResult(`Error: ${error.message}`, "error");
    });
}

function calculateLiquidityFund() {
    const liquidityFund = document.getElementById('liquidity_fund').value;
    const netWorth = document.getElementById('net_worth_liquidity').value;

    console.log('Input values:', {
        liquidityFund: liquidityFund,
        netWorth: netWorth
    });

    if (!liquidityFund || !netWorth) {
        displayResult("Please fill in all fields", "error");
        return;
    }

    // Create URLSearchParams for proper form encoding
    const formData = new URLSearchParams();
    formData.append('liquidity_fund', liquidityFund);
    formData.append('net_worth_liquidity', netWorth);

    console.log('Sending request with data:', formData.toString());

    // Use port 8081
    const fullUrl = 'http://localhost:8081' + API_ENDPOINTS['liquidity-ratio'];
    console.log('Full request URL:', fullUrl);

    fetch(fullUrl, {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: formData.toString()
    })
    .then(response => {
        console.log('Full response:', response);
        console.log('Response headers:', response.headers);
        console.log('Response status:', response.status);
        
        // Log the raw text of the response before parsing
        return response.text().then(text => {
            console.log('Raw response text:', text);
            try {
                return JSON.parse(text);
            } catch (error) {
                console.error('JSON parsing error:', error);
                throw new Error('Failed to parse JSON: ' + text);
            }
        });
    })
    .then(data => {
        console.log('Success response:', data);
        if (data.error) {
            displayResult(`Error: ${data.error}`, "error");
        } else {
            displayResult(`Liquidity Fund Ratio: ${data.percentage.toFixed(2)}%\n\n
                         <strong>Recommendation:</strong>\n
                         ${data.message}`, 
                         data.recommendation);
        }
    })
    .catch(error => {
        console.error('Error details:', error);
        displayResult(`Error: ${error.message}`, "error");
    });
}

// Function to navigate to specific calculator and show its form
function navigateToCalculator(calculatorType) {
    // First navigate to calculator page
    window.location.href = '/';
    
    // Store the calculator type in session storage
    sessionStorage.setItem('selectedCalculator', calculatorType);
}

// Function to show specific calculator form
function showCalculatorForm(formId) {
    // Hide all calculator forms
    const forms = document.querySelectorAll('.form-container');
    forms.forEach(form => {
        form.style.display = 'none';
    });
    
    // Show the selected form
    const selectedForm = document.getElementById(formId);
    if (selectedForm) {
        selectedForm.style.display = 'block';
        selectedForm.scrollIntoView({ behavior: 'smooth' });
        
        // Clear the result section
        const resultDisplay = document.getElementById('result');
        if (resultDisplay) {
            resultDisplay.innerHTML = '';
        }
    }
}

// Update existing showForm function to use the new showCalculatorForm
function showForm(formId) {
    showCalculatorForm(formId);
}

// Handle calculator navigation
function showCalculator(hash) {
    // Hide all calculator sections
    const calculators = document.querySelectorAll('.calculator-section');
    calculators.forEach(calc => calc.style.display = 'none');

    // Show the selected calculator
    const targetCalc = document.querySelector(hash);
    if (targetCalc) {
        targetCalc.style.display = 'block';
        targetCalc.scrollIntoView({ behavior: 'smooth' });
    }
}

// Listen for hash changes
window.addEventListener('hashchange', () => {
    if (window.location.hash) {
        showCalculator(window.location.hash);
    }
});

// Check hash on page load
window.addEventListener('load', () => {
    if (window.location.hash) {
        showCalculator(window.location.hash);
    }
});

// Function to show specific calculator based on hash
function showCalculatorFromHash() {
    const hash = window.location.hash.slice(1); // Remove the # symbol
    if (hash === 'savings') {
        showForm('savings-ratio');
        document.querySelector('[onclick="showForm(\'savings-ratio\')"]').classList.add('active');
    }
}

// Add event listeners
document.addEventListener('DOMContentLoaded', function() {
    // Show calculator based on hash when page loads
    showCalculatorFromHash();

    // Listen for hash changes
    window.addEventListener('hashchange', showCalculatorFromHash);
});