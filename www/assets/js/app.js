// import gsap from "gsap";

console.log(gsap.version)
document.querySelector('header').addEventListener('mousemove', moveHeader);

function moveHeader(e) {

    const { offsetX, offsetY, target } = e;
    const { clientWidth, clientHeight } = target;
    const xPos = (offsetX / clientWidth) - 0.5;
    const yPos  = (offsetY / clientHeight) - 0.5;
    const modifier = (index) => index * 1.2 * 0.5;

    const elements = gsap.utils.toArray('.header__primary__main, .header__primary__sub');

    elements.forEach((txt, index) => {
        gsap.to(txt,{
            duration: 1.5,
            x: xPos * 20 * modifier(index),
            y: yPos * 10 * modifier(index),
            rotationY: xPos * 40,
            rotateY: yPos * 10
        })
    })

}

function initNavigation() {}
function init(){}

window.addEventListener('load', function(){
    init();
});
