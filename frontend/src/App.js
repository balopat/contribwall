import React, {Component} from 'react';
import './App.css';

class App extends Component {
    constructor() {
        super();
        this.state = {
            contributors: [],
        }
    }

    componentDidMount() {
        fetch('/contributors')
            .then(results => {
                return results.json()
            })
            .then(data => {
                let contributors = data.map((contributor) => {
                    return (
                        <span>
                            <span >
                                <img alt={contributor.login} height="50px" width="50px" src={contributor.avatar}/>
                            </span>
                            <span>{contributor.login}</span>
                        </span>
                    )
                });
                this.setState({contributors: contributors})
            })
            .catch(error => {
                console.log(error)
            })
    }

    render() {
        return (
            <div className="App">
                <h1> Wall of contributors for minikube!</h1>
                {this.state.contributors}
            </div>
        );
    }
}

export default App;
