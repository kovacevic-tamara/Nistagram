import React, {useEffect, useState} from 'react';
import {Alert, Button, Dropdown, DropdownButton, FormControl, InputGroup, Card, Modal} from "react-bootstrap";
import axios from "axios";
import ProfileForSug from "../components/HomePage/ProfileForSug";
import {useHistory} from "react-router-dom";
import Navigation from "../components/HomePage/Navigation";
import searchService from "../services/search.service";
import {useSelector} from "react-redux";
import userService from "../services/user.service";
import Post from "../components/Post/Post";

import '../style/hover-post.css';
import '../style/Profile.css';
import { AiFillHeart } from 'react-icons/all';
import { FaComment } from 'react-icons/all';
import { IoIosHeartDislike } from 'react-icons/all';
import EditProfile from "../components/ProfileComponent/EditProfile";


export default function Search() {
    const [user,setUser] =useState({});
    const [searchCategory, setSearchCategory] = useState("Search category");
    const [input, setInput] = useState("");
    const [inputErr, setInputErr] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [searchResult, setSearchResult] = useState([]);
    const [searchPlaceholder, setSearchPlaceholder] = useState("search value");
    const store = useSelector(state => state);
    const [showModal, setModal] = useState(false);
    const [selectedPost, setSelectedPost] = useState({})


    useEffect(() => {
        //if(!props.location.state) window.location.replace("http://localhost:3000/unauthorized");

    },[])

    async function searchByUser() {
        const response = await userService.searchByUser({
            username: input,
            firstName: firstName,
            lastName: lastName,
            jwt: store.user.jwt,
        })

        if (response.status === 200)  setSearchResult(response.data.users);
        else  console.log("search error")
    }

    async function searchByTag() {
        const response = await searchService.searchByTag({
            text : input,
            jwt : store.user.jwt
        })

        if (response.status === 200)   setSearchResult(response.data.posts)
        else  console.log("search error")
    }

    async function searchByLocation() {
        const response = await searchService.searchByLocation({
            location : input,
            jwt : store.user.jwt
        })

        if (response.status === 200) setSearchResult(response.data.posts);
        else   console.log("search error loc")
    }

    function handleModal() {
        setModal(!showModal)
    }


    function search() {
        if (input === "" && searchCategory !== "User") {
            setInputErr("Enter search value")
            return
        }
        else if (input ==="" && firstName === "" && lastName === "" && searchCategory === "User") {
            setInputErr("Enter search value")
            return
        }

        if(!/^[a-zA-Z ,.'-]+$/.test(input) && searchCategory !== "User"){
            setInputErr("Enter valid search value")
            return
        }
        else if ( searchCategory === "User") {
            if( input!="" && !/^[a-zA-Z ,.'-]+$/.test(input)) {
                setInputErr("Enter valid search value")
                return
            }else  if( firstName!="" && !/^[a-zA-Z ,.'-]+$/.test(firstName)) {
                setInputErr("Enter valid search value")
                return
            }else  if( lastName!="" && !/^[a-zA-Z ,.'-]+$/.test(lastName)) {
                setInputErr("Enter valid search value")
                return
            }
        }

        switch (searchCategory) {
            case "User" :
                searchByUser();
                break;
            case "Tag" :
                searchByTag();
                break;
            case "Location" :
                searchByLocation();
                break;
            default:
                alert("Select search category");
                setInput("")
                return;
        }

        console.log(searchResult)
    }

    function handleSearchCategoryChange(event) {
        setSelectedPost({})
        setSearchResult([])
        setInput("")
        setFirstName("")
        setLastName("")
        setSearchCategory(event)
        switch (event) {
            case "Location" :
                setSearchPlaceholder("location");
                break;
            case "Tag" :
                setSearchPlaceholder("tag");
                break;
            case "User" :
                setSearchPlaceholder("username");
                break;
        }
    }

    function openPost(post) {
        setSelectedPost(post);
        handleModal();
    }


    return (
        <div  className="App">
            <Navigation d/>

            <br/>
            <br/><br/><br/><br/>
            <div className="row" style={{marginLeft : "10px"}}>
                <div className="col-sm-5 mb-2">
                    <DropdownButton onSelect={(e) => handleSearchCategoryChange(e) } as={InputGroup.Append}  variant="outline-secondary" title={searchCategory} id="input-group-dropdown-2" >
                        <Dropdown.Item eventKey={"Location"} >Location</Dropdown.Item>
                        <Dropdown.Item eventKey={"Tag"} >Tag</Dropdown.Item>
                        <Dropdown.Item eventKey={"User"} >User</Dropdown.Item>
                    </DropdownButton>
                </div>
                <div className="col-sm-5 mb-2" >
                    <input name="input" className="form-control" placeholder={searchPlaceholder} value={input} onClick={(e) => setInputErr("")} onChange={(e) => setInput(e.target.value)}/>
                    {inputErr.length > 0 &&
                    <span className="text-danger">{inputErr}</span>}
                </div>
                <div className="col-sm-5 mb-2" style={{display : "inline"}}>
                    {searchCategory === "User" &&
                    <input name="input" className="form-control" placeholder={"first name"} value={firstName} onClick={(e) => setInputErr("")} onChange={(e) => setFirstName(e.target.value)}/>}
                </div>
                <div className="col-sm-5 mb-2" style={{display : "inline"}}>
                    {searchCategory === "User" &&
                    <input name="input" className="form-control" placeholder={"last name"} value={lastName} onClick={(e) => setInputErr("")} onChange={(e) => setLastName(e.target.value)}/>}
                </div>

                <br/>
                <div className="col-sm-4">
                    <Button variant="primary" onClick={search}>Search</Button>{' '}
                </div>
            </div>
            <br/><br/>

            {searchResult.length > 0 && searchCategory === 'User' &&
                <ul style={{marginLeft : "30px"}}>
                    {searchResult.map((user, i) =>
                        <ProfileForSug user={user} username={user.username} firstName={user.firstName} lastName={user.lastName} caption={user.biography} urlText="Follow" iconSize="big" captionSize="small"  storyBorder={true} />
                    )}
                </ul>
            }


            <div  style={{marginLeft: '20%', marginRight: '20%',marginTop:'10%'}}>
                <div className="gallery">


                    {searchResult.length > 0 && searchCategory !== 'User' &&
                        searchResult.map((post) => {
                            return (
                                <div className="content" onClick={() => openPost(post)}>
                                    <div className="doc" >
                                        <img src={post.media[0].content} alt="" className="item"/>
                                        <div className="links">
                                            <a href=""><i >{post.likes.length}<AiFillHeart /></i></a>
                                            <a href=""><i >{post.dislikes.length}<IoIosHeartDislike /></i></a>
                                            <a href=""><i >{post.comments.length}<FaComment /></i></a>
                                        </div>
                                    </div>
                                </div>
                            );
                        })
                    }
                </div>
            </div>

            <Modal show={showModal} onHide={handleModal}>
                <Modal.Body>
                    <Post post={selectedPost} postUser={{ id: selectedPost.userId }}/>
                </Modal.Body>
            </Modal>
        </div>
    );
}