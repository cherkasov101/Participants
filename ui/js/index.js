const form = document.getElementById('registration-form');
const nameInput = document.getElementById('name');
const surnameInput = document.getElementById('surname');
const patronymicInput = document.getElementById('patronymic');
const phoneInput = document.getElementById('phone');
const emailInput = document.getElementById('email');

form.addEventListener('submit', function(event) {
    if (!nameInput.checkValidity()) {
        alert('Пожалуйста, введите корректное имя');
        event.preventDefault();

        return;
    }

    if (!surnameInput.checkValidity()) {
        alert('Пожалуйста, введите корректную фамилию');
        event.preventDefault();
        return;
    }

    if (!patronymicInput.checkValidity()) {
        alert('Пожалуйста, введите корректное отчество');
        event.preventDefault();
        return;
    }

    if (!phoneInput.checkValidity()) {
        alert('Пожалуйста, введите корректный телефон');
        event.preventDefault();
        return;
    }

    if (!emailInput.checkValidity()) {
        alert('Пожалуйста, введите корректный email');
        event.preventDefault();
        return;
    }
});