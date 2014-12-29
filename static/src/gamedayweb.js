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
        var data = this.props.data;
        if (typeof data.game !== 'undefined') {
            return (
                <div className="gameBoxscore">
                    <h2>{data.game.AwayTeamName} at {data.game.HomeTeamName}</h2>
                    <h3>{data.game.Venue} - {data.game.Status}</h3>
                    <p>{data.game.Gameday}</p>
                    <p>{data.linescore.HomeTeamRuns}</p>
                    <p>{data.linescore.AwayTeamRuns}</p>
                </div>
                );
        }
        return (<div/>);
    }
});

React.render(
    <GameBox url="game" />,
    document.getElementById('content')
);
