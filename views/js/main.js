document.querySelectorAll('.nav-link').forEach(link => {
    link.addEventListener('click', function () {
        document.querySelectorAll('.nav-link').forEach(nav => {
            nav.classList.remove('active');
            nav.setAttribute('aria-selected', 'false');
        });
        this.classList.add('active');
        this.setAttribute('aria-selected', 'true');
    });
});

const buttonQuery = document.getElementById('btnQuery');
// buttonQuery.style.display = "none";

function fetchUsers() {
    fetch('/users')
        .then(response => response.json())
        .then(data => {
            let resultDiv = document.getElementById('result');
            resultDiv.innerHTML = '';
            if (data.length > 0) {
                let table = '<table class="table" border="1"><tr><th>ID</th><th>Username</th><th>Password</th><th>Role</th></tr>';
                data.forEach(user => {
                    table += `<tr><td>${user.id}</td><td>${user.username}</td><td>${user.password}</td><td>${user.role}</td></tr>`;
                });
                table += '</table>';
                resultDiv.innerHTML = table;
            } else {
                resultDiv.innerHTML = 'No users found';
            }
        })
        .catch(error => console.error('Error fetching users:', error));
}


function bypassIndex() {
    location.replace('/index');
}

// function showFeatures() {
//     let ariaTrue = document.getElementById('btnTest').classList.contains('active');
//     // let ariaFalse = buttonQuery.getAttribute('aria-selected', 'false');

//     if (ariaTrue === true) {
//         buttonQuery.style.display = 'block';
//     } else {
//         buttonQuery.style.display = 'none';
//     }
//     return
// }



// document.getElementById('btnQuery').addEventListener("click", function () {
//     alert("Success");
// })

