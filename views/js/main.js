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
document.getElementById('btnQuery').addEventListener("click", function () {
    alert("Success");
})
