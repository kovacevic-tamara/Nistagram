import React from 'react';
import { Modal } from 'react-bootstrap';
import Post from './Post';
import "./../../style/PostPreviewModal.css"

const PostPreviewModal = (props) => {
    const { post, showModal, setShowModal, shouldReload, isAd, setPosts } = props;

    return (
        <Modal 
            className="PostPreviewModal__Wrapper"
            contentClassName="content" 
            show={showModal} 
            onHide={() => setShowModal(false)}>
            <Post className="Post" shouldReload={shouldReload} post={post} setPosts={setPosts} isAd={isAd} />
        </Modal>
    )
}

export default PostPreviewModal;