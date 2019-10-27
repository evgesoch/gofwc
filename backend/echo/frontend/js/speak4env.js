$(function() {
  listenAndActOnSavePostButtonPress();

  var myObject3 = new Vue({
    el: '#app-3',
    data: {message: 'Hello GalZ!'}
  });

});

/**
 * Respond to clicks on Save Post button in the modal
 */
function listenAndActOnSavePostButtonPress() {
  $("#savePostButton").on('click', () => {
    let postElem       = $("#post");
    let postsContainer = $("#postsContainer");
    let postText       = $("#postText").html();
    let data           = prepareNewNoteData(postText);

    $.when(makeAjaxRequest("GET", "http://localhost:8080/posts/1", JSON.stringify(data)))
      .done((response) => {
        console.log(response);
        $("#postText").html("");
        $("#successMessage").fadeIn().removeClass("d-none").delay(2000).fadeOut();
        
      })
      .fail((response) => {
        $("#errorMessage").fadeIn().removeClass("d-none").delay(2000).fadeOut();


      })
      .always((reponse) => {

      });



  });
}

/**
 * Prepare data for creating a new Note
 * 
 * @param {String}  postText new post's text
 * 
 * @return {Object}          JS object with payload data
 */
function prepareNewNoteData(postText) {
  return {
    text: postText
  }
}

/**
 * Wrapper for jQuery $.ajax
 * 
 * @param {String} method The HTTP method
 * @param {String} url    The url in which the request is sent
 * @param {JSON}   data   The payload data
 * 
 * @return {jqXHR}        jQuery XHR object
 */
function makeAjaxRequest(method, url, data) {
  return (
    $.ajax({
      method: method,
      url: url,
      data: data,
      contentType: "application/json"
    })
  );
}
