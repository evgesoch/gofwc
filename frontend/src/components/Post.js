import React from 'react';
import '../css/Post.css';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';

function Post(props) {
    return (
        <>
        <Container style={{paddingRight: 100, paddingLeft: 100, marginBottom: 15}}>
            <Card>
                <Card.Header className={"s4e-postHeader"} style={{backgroundColor: "#faeaa2", padding: "5px 20px"}}>Post #{props.postID}</Card.Header>
                <Card.Body>
                    <Card.Text>
                        {props.postText}
                    </Card.Text>
                </Card.Body>
            </Card>
        </Container>
        </>
    );
}

export default Post;
