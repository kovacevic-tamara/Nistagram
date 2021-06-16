import "../../style/menu.css";
import { ReactComponent as Home } from "../../images/home.svg";
import { ReactComponent as Inbox } from "../../images/inbox.svg";
import { ReactComponent as Notifications } from "../../images/notifications.svg";
import { ReactComponent as Bookmark } from "../../images/bookmark.svg";
import { ReactComponent as Plus } from "../../images/plus.svg";
import { ReactComponent as Explore } from "../../images/more.svg";
import { ReactComponent as VerificationSymbol } from "../../images/verification-symbol.svg";

import ProfileIcon from "../ProfileComponent/ProfileIcon";
import {NavLink, Route} from "react-router-dom";
import Profile from "../ProfileComponent/Profile";
import {useEffect, useState} from "react";
import userService from "../../services/user.service";
import {useDispatch, useSelector} from "react-redux";
import {Dropdown} from "react-bootstrap";

function Menu() {
    const[username,setUsername]=useState('')
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        //   if(!props.location.state) window.location.replace("http://localhost:3000/unauthorized");
        getUsernameById();
    }, []);

    async function getUsernameById() {
        const response = await userService.getUsernameById({
            id: store.user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUsername(response.data.username)
        } else {
            console.log("getuserbyusername error")
        }
    }
    return (
        <div className="menu">
            <NavLink  to={{ pathname: "/home" }}  >
                <Home className="icon"/>
            </NavLink>
            <NavLink  to={{ pathname: "/chats" }}  >
                <Inbox className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/notifications" }}  >
                <Notifications className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/saved" }}  >
                <Bookmark className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/newpost" }}  >
                <Plus className="icon" />
            </NavLink>
            {/*<NavLink  to={{ pathname: "/submit-verification-request" }}  >*/}
                <Dropdown>
                    <Dropdown.Toggle variant="link" id="dropdown-basic">
                        <VerificationSymbol className="icon" />
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                        <Dropdown.Item href="/submit-verification-request">Submit verification request</Dropdown.Item>
                        <Dropdown.Item href="/view-my-verification-request">View my verification requests</Dropdown.Item>
                    </Dropdown.Menu>
                </Dropdown>
            {/*</NavLink>*/}

            <NavLink  to={"/profile/"+username }  >
                <ProfileIcon iconSize="medium" image='https://i.pravatar.cc/150?img=1' />
            </NavLink>

            <NavLink  to={{ pathname: "/info" }} >
                <Explore className="icon" />
            </NavLink>

            <a href='/' style={{marginLeft:'10px'}}>Log out</a>

        </div>
    );
}

export default Menu;