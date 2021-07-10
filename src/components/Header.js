import React from 'react';

const Header = (props) => (
    <div className="m-0 p-0">
        <header className="container">
            <div className="row">
                <h1><a href="/#/home" id="Logo">Logo</a></h1>
                <nav id="hrefpNav" className="text-right">
                    <a href="/#/home">Home</a>
                    <a href="/#/about">About</a>
                </nav>
            </div>
        </header>
        <section className="jumbotron">
            <div className="container">
                <div className="row">
                    <h2>{props.page}</h2>
                </div>
            </div>
        </section>
    </div>
);

export default Header;