a page

/*
1. Refresh page with an animation
*/
 
setInterval(function(){
 
 //get the current page url
 var currentUrl = window.location.href;
 
 //create a fade out animation for the page
 $("#main").fadeOut(2500);
 
 //redirect to the same page after the animation has finished
 setTimeout(function(){
   window.location.href = currentUrl;
 }, 2500);
 
}, 10000);


/*
2. Refresh page with a notification
*/

//set interval to refresh page every 10 seconds
setInterval(function(){

//display refresh notification
alert("Refreshing page...");

//get the current page url
var currentUrl = window.location.href;

//redirect to the same page
window.location.href = currentUrl;

}, 10000);


/*
3. Refresh page with a message
*/

//set interval to refresh page every 10 seconds
setInterval(function(){

//display refresh message in the page
$("#main").html('<div class="refresh-message">Refreshing page...</div>');

//get the current page url
var currentUrl = window.location.href;

//redirect to the same page
window.location.href = currentUrl;

}, 10000);


/*
4. Refresh page with a counter
*/

//set interval to refresh page every 10 seconds
var sec = 10;
setInterval(function(){

//display refresh counter in the page
$("#main").html('<div class="refresh-counter">Refreshing page in ' + sec + ' seconds...</div>');

//decrement counter
sec--;

//check if counter has reached 0
if (sec == 0) {

//get the current page url
var currentUrl = window.location.href;

//redirect to the same page
window.location.href = currentUrl;

//reset the counter
sec = 10;

}

}, 1000);


/*
5. Refresh page with a sound
*/

//set interval to refresh page every 10 seconds
setInterval(function(){

//play a sound to notify the user
var sound = new Audio('refresh-sound.mp3');
sound.play();

//get the current page url
var currentUrl = window.location.href;

//redirect to the same page
window.location.href = currentUrl;

}, 10000);