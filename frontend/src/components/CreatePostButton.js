import React, {useState, createElement} from 'react';
import Post from './Post';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import Container from 'react-bootstrap/Container';
import Alert from 'react-bootstrap/Alert'
import '../css/CreatePostButton.css';
import axios from 'axios';

function CreatePostButton(props) {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    function createPost() {
        let postText = document.getElementById("postText").innerHTML;
                console.log(postText);
        //let newPost = createElement(Post, postText);
        handleClose();


        
        /*let successMessageElem = document.getElementById("successMessage");
        let errorMessageElem = document.getElementById("errorMessage");

        // An pane ola kala me to request deixe success message kai kleise error
        // message an einai anoixto
        // Meta apo 3 sec kleise modal kai success message
        errorMessageElem.classList.add("d-none");
        successMessageElem.classList.remove("d-none");
        setTimeout(() => {
            handleClose();
            successMessageElem.classList.add("d-none");
        }, 3000);
        
        // Alliws mhn kleineis to modal kai deixe error messsage mexri na petyxei
        // h dhmiourgia kaniouriou Post
        errorMessageElem.classList.remove("d-none");*/

    }
    
    return (
        <>
        <Container className={"text-center"}>
            <Button variant={"success"} className={"rounded-pill"} onClick={handleShow}>
                + Create Post
            </Button>
        </Container>

        <Modal show={show} onHide={handleClose}>
            <Alert variant="success" className={"d-none"} id={"successMessage"}>
                <Alert.Heading>
                    Success
                </Alert.Heading>
                The Post was published successfully!
            </Alert>
            <Alert variant="danger" className={"d-none"} id={"errorMessage"}>
                <Alert.Heading>
                    Error
                </Alert.Heading>
                An error occurred when publishing the Post. Please try again.
            </Alert>
            <Modal.Header closeButton>
                <Modal.Title>Create a new Post</Modal.Title>
            </Modal.Header>
            <Modal.Body className={"p-0"}>
                <div className={"postText"} id={"postText"} contentEditable={"true"}></div>
            </Modal.Body>
            <Modal.Footer>
                <Button variant={"outline-secondary"} onClick={handleClose}>
                    Close
                </Button>
                <Button variant={"outline-success"} onClick={createPost}>
                    Save Post
                </Button>
            </Modal.Footer>
        </Modal>
        </>
    );
}

export default CreatePostButton;
