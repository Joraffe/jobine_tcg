export const renderHtmlTo = (selector, htmlString) => {
  const element = document.querySelector(selector);

  element.innerHtml = htmlString;
};


export const clearHtmlFor = (selector) => {
  const element = document.querySelector(selector);
  
  element.innerHtml = '';
};


export const addClassTo = (selector, classes) => {
  const element = document.querySelector(selector);
  element.classList.add(...classes);
};


export const clearClassesFor = (selector) => {
  const element = document.querySelector(selector);
  element.classList.remove( )
};
