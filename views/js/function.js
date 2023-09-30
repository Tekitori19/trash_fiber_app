export function setCurrentCss(path) {
    const currentCSS = "bg-indigo-700 text-white".split(" ")
    const defaultCSS = "text-indigo-200 hover:text-white hover:bg-indigo-700".split(" ")
    
    let node = document.querySelectorAll("nav a")
    
    for (let index = 0; index < node.length; index++) {
        let Element = node[index];
    
        if (Element.innerText == path) {
            Element.classList.add(...currentCSS)
        }else {
            Element.classList.remove(...currentCSS)
            Element.classList.add(...defaultCSS)
        }
        
    }
}
