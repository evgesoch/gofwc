import React from 'react';
import tree from '../images/tree.svg';
import '../css/Header.css';

class Header extends React.Component {
    render() {
        return (
            <div className={"container"}>               
                <div className={"jumbotron bg-white mb-10"}>
                    <div className={"container text-center"}>
                        <div className={"d-inline text-center"}>
                            <img src={tree} className={"treeLogo"} alt={"Speak4Env logo"}></img>
                        </div>
                        <div className={"d-inline s4e-title h1"}>
                            Speak4Env!
                        </div>
                    </div>
                    <hr className={"my-4"}/>
                    <p className={"lead"}><span className={"font-weight-bold"}>Speak4Env</span> gives you the opportunity to express and share with others your thoughts and ideas about improving our planet's environment. Let's make it a better place for all of us!</p>
                </div>
            </div>
        );
    }
}

export default Header;
