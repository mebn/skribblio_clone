function _(selector){
  return document.querySelector(selector);
}
function setup(){
  let canvas = createCanvas(750, 700);
  canvas.parent("canvas-wrapper");
  background(255);
  strokeWeight(10); 
}
function resize(clicked_id){
  let size = 0; 
  if(clicked_id == "one") size = 10; 
  else if (clicked_id == "two") size = 20; 
  else if (clicked_id == "three") size = 30; 
  else size = 40; 
  strokeWeight(size); 
}
function mouseDragged(){
  let type = _("#pen-brush").checked?"brush":"eraser";
  let color = _("#color-picker").value;
  stroke(color);
  if(type == "brush"){
    line(pmouseX, pmouseY, mouseX, mouseY);
  } else {
    stroke(255);
    line(pmouseX, pmouseY, mouseX, mouseY);
  }
}
_("#reset-canvas").addEventListener("click", function(){
  background(255);
});