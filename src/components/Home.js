import React from 'react';
import { datas } from '../data/data';

const people = [
  "Siri",
  "Alexa",
  "Google",
  "Facebook",
  "Twitter",
  "Linkedin",
  "Sinkedin"
];

function Home () {
    const [searchTerm, setSearchTerm] = React.useState("");
    const [searchResults, setSearchResults] = React.useState([]);
    const handleChange = event => {
        setSearchTerm(event.target.value);
    };
    React.useEffect(() => {
        // console.log("------HERE datas------",datas)
        const results = datas.map(data => {
            console.log(data['title'].toLowerCase().includes(searchTerm))
            return (data['title'].toLowerCase().includes(searchTerm) ? data : false)
        })
        console.log(results)
        setSearchResults(results);
    }, [searchTerm]);

    const cardBody = (title, href, scr) => (
        <div className='card m-1 p-2'>
            <a className="p-0" href={href}>
                <img className="card-img-top" src={scr} alt={title}/>
            </a>
            <div className="card-body">
                <h5 className="card-title" id="card-text" style={{color:'black'}}>{title}</h5>
            </div>
        </div>
    )


    return (
        <div className="m-0 p-0">
            <section className="jumbotron">
                <div className="container">
                    <div className="row">
                        <h2 className="text-success">Vegan Recipes!</h2>
                    </div>

                    <input 
                        className="form-control mr-sm-2 col-12 mb-1" 
                        type="search" 
                        placeholder="Search" 
                        aria-label="Search" 
                        value={searchTerm}
                        onChange={handleChange}
                    /> 
                        
                </div>
            </section>
    
            <section className="container">

                 <div className="card-columns m-1">
                    {searchResults.map(item => (
                        item ? cardBody(item['title'], item['href'], item['img']) : ""
                    ))}
                </div>
                
            </section>
            
        </div>
    );
}

export default Home;