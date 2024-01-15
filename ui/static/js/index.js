const form = document.getElementById('registration-form');
const nameInput = document.getElementById('name');
const surnameInput = document.getElementById('surname');
const patronymicInput = document.getElementById('patronymic');
const birthdateInput = document.getElementById('birthdate');
const sectionInput = document.getElementById('section');
const phoneInput = document.getElementById('phone');
const emailInput = document.getElementById('email');
const presentationCheckbox = document.getElementById('presentation');
const topicContainer = document.getElementById('topic-container');
const topicInput = document.getElementById('topic');

presentationCheckbox.addEventListener('change', function() {
    if (presentationCheckbox.checked) {
        topicContainer.style.display = 'block';
        topicInput.required = true;
    } else {
        topicContainer.style.display = 'none';
        topicInput.required = false;
    }
});

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

    const formData = {
        name: nameInput.value,
        surname: surnameInput.value,
        patronymic: patronymicInput.value,
        birthdate: birthdateInput.value,
        section: sectionInput.value,
        phone: phoneInput.value,
        email: emailInput.value,
        presentation: presentationCheckbox.checked ? 'yes' : 'no',
        topic: topicInput.value
    };

    const jsonData = JSON.stringify(formData);

    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/send');
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                console.log(xhr.responseText);
            } else {
                console.error('Произошла ошибка:', xhr.status);
            }
        }
    };
    xhr.send(jsonData);
    window.location.href = "/answers";
});