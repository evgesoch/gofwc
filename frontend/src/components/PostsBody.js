import React, {useState} from 'react';
import Container from 'react-bootstrap/Container';
import Post from './Post'

function PostsBody(props) {
    const [numPosts, setNumPosts] = useState(0);
    const children = [];

    // Get the number of posts from the props and assign them to the state
    var getNumPosts = () => setNumPosts(props.numChildrenPosts);
 

    function addPosts() {
        for (let i = 0; i < numPosts; i++) {
            children.push(<Post postID={0} postText="hello" />);
        }
    }
        

    return (
        <div id="postsBody">
            <Post postID="6" postText="This is a newer post"/>
            <Post postID="5" postText="This is a new post"/>
        </div>
    );
}

export default PostsBody;
