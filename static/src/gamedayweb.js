var GameBox = React.createClass({
    getInitialState: function() {
        return { data: {} };
    },
    componentDidMount: function() {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            success: function(data) {
                this.setState({data: data});
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    render: function() {
        return (
            <div className="gameBox">
                <GameForm/>
                <GameBoxscore data={this.state.data}></GameBoxscore>
            </div>
        );
    }
});

var GameForm = React.createClass({
    render: function() {
        return (
            <div className="gameForm"/>
        );
    }
});

var GameBoxscore = React.createClass({
    render: function() {
        var game = this.props.data;
        return (
            <div className="gameBoxscore">
                <h2>{game.AwayTeamName} at {game.HomeTeamName}</h2>
                <h3>{game.Venue} - {game.Status}</h3>
                <p>{game.Gameday}</p>
            </div>
            );
    }
});

React.render(
    <GameBox url="game" />,
    document.getElementById('content')
);
