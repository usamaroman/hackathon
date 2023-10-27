import "./taskcard.css"

export const TaskCard = (props) => {
    const {title, description, start, end, difficulty} = props.pr

    return (
        <div class="container">
            <div class="card">
                <div class="info">
                    <div class="shoeName">
                        <div>
                            <h1 class="big">{title}</h1>
                        </div>
                    </div>
                    <div class="description">
                        <h3 class="title">Описание</h3>
                        <p class="text">{description}</p>
                    </div>
                    <div class="buy-price">
                        <div class="price">
                            <i class="fas fa-dollar-sign"></i>
                            <h5>{start} - {end}</h5>
                        </div>
                    </div>
                    <div class="buy-price">
                        <div class="price">
                            <i class="fas fa-dollar-sign"></i>
                            <h3 class="title">Сложность</h3>
                            <h5>{difficulty}</h5>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}