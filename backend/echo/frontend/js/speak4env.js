$(function() {
  listenAndActOnSavePostButtonPress();

  var myObject3 = new Vue({
    el: '#app-3',
    data: {message: 'Hello GalZ!'}
  });

});

/**
 * Fetch all Posts from the database and display them
 */
function fetchAndRenderAllPosts() {

}

/**
 * Respond to clicks on Save Post button in the modal and create a new Post
 */
function listenAndActOnSavePostButtonPress() {
  $("#savePostButton").on('click', () => {
    let postText = $("#postText").html();
    let data     = prepareNewNoteData(postText);

    $.when(makeAjaxRequest("POST", "http://localhost:8080/posts", JSON.stringify(data)))
      .done((response) => {
        let clonedPostElem = $("#post").clone().removeClass("d-none");
        let firstPost      = $("#postsContainer").children().first();

        $("#postText").html("");
        $("#successMessage").fadeIn().removeClass("d-none").delay(2000).fadeOut();
        clonedPostElem.find(".s4e-postHeader").html(`Post #${response.postID}`);
        clonedPostElem.find("p").html(postText);
        clonedPostElem.insertBefore(firstPost);
      })
      .fail((response) => {
        $("#errorMessage").fadeIn().removeClass("d-none").delay(2000).fadeOut();
      })
  });
}

/**
 * Prepare data for creating a new Post
 * 
 * @param {String} postText The new Post's text
 * 
 * @return {Object}         JS object with payload data
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
