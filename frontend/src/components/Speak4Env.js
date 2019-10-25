import React from 'react';
import '../css/Speak4Env.css';
import Header from './Header';
import CreatePostButton from './CreatePostButton';
import PostsBody from './PostsBody';
import Footer from './Footer';

class Speak4Env extends React.Component {
    render() {
        return (
            <React.Fragment>
                <Header/>
                <CreatePostButton/>
                <PostsBody/>
                <Footer/>
            </React.Fragment>
        );
    }
}

export default Speak4Env;
