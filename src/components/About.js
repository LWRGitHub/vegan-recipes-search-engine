import React from 'react';

import Header from './Header';

const MainContent = () => (
    <section className="container">
      <div className="row justify-content-center">
        <div className="col-sm-6">

          {/* About info & img */}
          {/* <img className='rounded' src="..." alt="..."/> */}
          <p>Hello,</p>
          <p>We are here to provide you with the vegan search engine that you have dream of...</p>
          <p>...</p>
          <p>...</p>
          {/* About info & img */}


        </div>
      </div>
    </section>
)

const About = () => (
    <div>
        <Header page={'About Me'} />
        <MainContent />
    </div>
);

export default About;