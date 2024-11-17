// Fetch all invoices and display them in the table
function loadInvoices() {
    fetch('/api/invoices')
        .then(response => response.json())
        .then(invoices => {
            const tableBody = document.getElementById('invoice-table').getElementsByTagName('tbody')[0];
            tableBody.innerHTML = ''; // Clear existing rows

            invoices.forEach(invoice => {
                const row = tableBody.insertRow();
                row.innerHTML = `
                    <td>${invoice.id}</td>
                    <td>${invoice.customer}</td>
                    <td>$${invoice.amount.toFixed(2)}</td>
                    <td>${new Date(invoice.due_date).toLocaleDateString()}</td>
                    <td>${invoice.description}</td>
                `;
            });
        })
        .catch(error => console.error('Error loading invoices:', error));
}

// Handle form submission to create a new invoice
document.getElementById('invoice-form').addEventListener('submit', function(event) {
    event.preventDefault();

    const customer = document.getElementById('customer').value;
    const amount = parseFloat(document.getElementById('amount').value);
    const dueDate = document.getElementById('dueDate').value;
    const description = document.getElementById('description').value;

    const invoiceData = {
        customer,
        amount,
        due_date: dueDate,
        description
    };

    fetch('/api/invoice', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(invoiceData)
    })
    .then(response => response.json())
    .then(invoice => {
        console.log('Invoice created:', invoice);
        loadInvoices(); // Reload the invoices list
    })
    .catch(error => console.error('Error creating invoice:', error));
});

// Load invoices when the page is loaded
window.onload = loadInvoices;
