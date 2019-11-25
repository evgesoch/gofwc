fetchAndRenderAllPosts();

$(function () {
	listenAndActOnSavePostButtonPress();
	scrollPageToTop();
	fetchAllPostsEvery5Sec();
});

/**
 * Fetch all Posts from the database and display them
 */
function fetchAndRenderAllPosts() {
	let postsContainer = $("#postsContainer");

	$.when(makeAjaxRequest("GET", "http://localhost:8080/posts"))
	.done((response) => {
		postsContainer.empty();

		for (let i = 0; i < response.length; i++) {
			let clonedPostElem = $("#post").clone();

			clonedPostElem.find(".s4e-postHeader").html(`Post #${response[i].ID}`);
			clonedPostElem.find("p").html(response[i].Text);
			postsContainer.append(clonedPostElem);
			clonedPostElem.removeClass("d-none");
		}
	})
	.fail((response) => {
		$("#errorAllPostsMessage").fadeIn().removeClass("d-none").delay(2000).fadeOut();
	})
}

/**
 * Respond to clicks on Save Post button in the modal and create a new Post
 */
function listenAndActOnSavePostButtonPress() {
	$("#savePostButton").click((e) => {
		let postText = $("#postText").html();
		let data = prepareNewPostData(postText);

		$.when(makeAjaxRequest("POST", "http://localhost:8080/posts", JSON.stringify(data)))
		.done((response) => {
			let clonedPostElem = $("#post").clone().removeClass("d-none");
			let firstPost = $("#postsContainer").children().first();

			$("#postText").html("");
			$("#successMessage").fadeIn().removeClass("d-none").delay(2000).fadeOut();
			clonedPostElem.find(".s4e-postHeader").html(`Post #${response.postID}`);
			clonedPostElem.find("p").html(postText);
			clonedPostElem.insertBefore(firstPost).hide().fadeIn();
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
function prepareNewPostData(postText) {
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

/**
 * Refresh the Posts every 5 sec
 */
function fetchAllPostsEvery5Sec() {
	setInterval(() => fetchAndRenderAllPosts(), 5000);
}

/**
 * Scroll the page to top
 */
function scrollPageToTop() {
	$(window).scroll((e) => {
		if ($(this).scrollTop() > 50) {
			$('#backToTop').fadeIn();
		} else {
			$('#backToTop').fadeOut();
		}
	});

	$('#backToTop').click((e) => {
		e.preventDefault();

		$('body,html').animate({
			scrollTop: 0
		}, 400);
	});
}
