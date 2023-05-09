const menu = document.querySelector('.burger');
const burgerIco = document.querySelector('.burger__wrapper');
const closeBurg = document.querySelector('.burger__close')

function openBurger() {
	menu.style.transform = "translateX(0%)"
    burgerIco.style.display = 'none'
	closeBurg.style.display = 'flex';
	setTimeout(() => {
		menu.style.display = 'flex';
	}, 10)
}
function closeBurger() {
	burgerIco.style.display = 'flex';
	closeBurg.style.display = 'flex';
	menu.style.transform = "translateX(100%)";
	setTimeout(() => {
		menu.style.display = 'none';
	}, 250)
}
