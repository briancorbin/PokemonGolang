class PokeApp extends React.Component {
    render() {
        return (
            <div className="pokeapp">
                <h1> The Kanto Pokedex! </h1>
                <PokemonList/>
            </div>
        );
    }
}

class PokemonList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            species: [],
            fetched: false,
            loading: false,
        };
    }
    componentWillMount() {
        this.setState({
            loading: true
        });

        fetch('http://pokeapi.co/api/v2/pokemon?limit=151').then(res => res.json()).then(response => {
            this.setState({
                species: response.results,
                loading: true,
                fetched: true,
            });
        });
    }
    render() {
        const {fetched, loading, species} = this.state;
        let content;
        if (fetched) {
            content = (
                <div className="pokemon--species--list">
                {species.map((pokemon,index)=><Pokemon key={pokemon.name} id={index+1} pokemon={pokemon}/>)}
                </div>
            );
        } else if (loading && !fetched) {
            content = <p> Loading... </p>;
        } else {
            content = <div></div>;
        }

        return <div>{content}</div>;
    }
}

class Pokemon extends React.Component {
    render() {
        const {pokemon, id} = this.props;
        return (
            <div className="pokemon--species">
                <div className="pokemon--species--container">
                    <div className="pokemon--species--sprite">
                        <img src={`/public/img/sprites/${id}.png`} />
                    </div>
                    <div className="pokemon--species--name"> {pokemon.name} </div>
                </div>
            </div>
        )
    }
}

ReactDOM.render(<PokeApp/>, document.getElementById('app'));