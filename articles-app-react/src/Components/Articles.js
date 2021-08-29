import React, { Component, useEffect, useState} from "react";

function Articles() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);
  
    // Remarque : le tableau vide de dépendances [] indique
    // que useEffect ne s’exécutera qu’une fois, un peu comme
    // componentDidMount()
    useEffect(() => {
      fetch("http://localhost:10000/articles")
        .then(res => res.json())
        .then(
          (result) => {
            setIsLoaded(true);
            setItems(result);
            console.log(result)
          },
          // Remarque : il faut gérer les erreurs ici plutôt que dans
          // un bloc catch() afin que nous n’avalions pas les exceptions
          // dues à de véritables bugs dans les composants.
          (error) => {
            setIsLoaded(true);
            setError(error);
          }
        )
    }, [])

    function deleteArticle(id) {
      // Simple DELETE request fetch
      const requestOptions = {
          method: 'DELETE'
      };
      fetch("http://localhost:10000/article/" + id, requestOptions)
      .then(res => res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          setItems(result);
          console.log(result)
        },
        // Remarque : il faut gérer les erreurs ici plutôt que dans
        // un bloc catch() afin que nous n’avalions pas les exceptions
        // dues à de véritables bugs dans les composants.
        (error) => {
          setIsLoaded(true);
          setError(error);
        }
      )
    }

    if (error) {
      return <div>Erreur : {error.message}</div>;
    } else if (!isLoaded) {
      return <div>Chargement...</div>;
    } else {
      return (
        <table className="table table-hover">
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Description</th>
                    <th>Content</th>
                    <th></th>
                    <th></th>
                </tr>
            </thead>
            {items!=null?
            <tbody>
            {items.map(item => (
            <tr key={item.Id}>
                    <td>{item.Title}</td>
                    <td>{item.Desc}</td>
                    <td>{item.Content}</td>
                    <td><input type="button" className="btn btn-warning" value="Modifier"/></td>
                    <td><input type="button" className="btn btn-danger" value="Supprimer" onClick={() => deleteArticle(item.Id)}/></td>
            </tr>
          ))}
            </tbody>:<tr></tr>}
        </table>
      );
    }
  }

export default Articles;