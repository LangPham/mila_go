import "../css/app.sass";

"use strict";

document.addEventListener('DOMContentLoaded', () => {
    window.addEventListener("click", function (e) {
        var element = e.target;
        let thetaLinkEvent = new CustomEvent('theta.link.click', {
            "bubbles": true, "cancelable": true
        });
        if (!element.dispatchEvent(thetaLinkEvent)) {
            console.log(element.getAttribute("data-method") === "get")
            e.preventDefault();
            e.stopImmediatePropagation();
            return false;
        }

        if (element.getAttribute("data-method")) {
            handleClick(element);
            e.preventDefault();
            return false;
        } else {
            element = element.parentNode;
        }
    }, false)

    function handleClick(element) {

        let to = element.getAttribute("href"),
            method = buildHiddenInput("_METHOD", element.getAttribute("data-method")),
            form = document.createElement("form");

        form.method = (element.getAttribute("data-method") === "get") ? "get" : "post";
        form.action = to;
        form.style.display = "hidden";
        form.appendChild(method);
        document.body.appendChild(form);
        form.submit();
    }

    function buildHiddenInput(name, value) {
        var input = document.createElement("input");
        input.type = "hidden";
        input.name = name;
        input.value = value;
        return input;
    }

    window.addEventListener('theta.link.click', function (e) {
        var message = e.target.getAttribute("data-confirm");
        if (message && !window.confirm(message)) {
            e.preventDefault();
        }
    }, false);
})
