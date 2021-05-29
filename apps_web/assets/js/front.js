import "../css/app.sass";
document.addEventListener('DOMContentLoaded', () => {

	let link = document.querySelectorAll('a[data-method]');
	if (link.length >0) {
		for (let i = 0; i < link.length; i++) {
			link[i].style.color = "red";
			link[i].addEventListener("click", function (e){
				let txt;
				 let r = confirm("DELETE!");
				if (r === true) {
					txt = "You pressed OK!";
				} else {
					txt = "You pressed Cancel!";
				}
				console.log(txt);
			})
		}
	}
})


