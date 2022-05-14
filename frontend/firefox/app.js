player = document.wrappedJSObject.getElementById("movie_player");
setInterval(function(){
  vidid = document.location.href.split("v=")[1].substring(0, 11)
  var t = ~~player.getCurrentTime();
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function(){
      if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
          sendResponse();
      }
  };
  xhr.open("GET", 'http://localhost:3000/video?vidid='+vidid+'&t='+t, true);
  xhr.send();
}, 1000);