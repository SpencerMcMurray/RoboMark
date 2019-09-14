import React, {useState} from 'react';
import { Button } from 'react-bootstrap';

const TestButtons = props => {
    const [count, setCount] = useState(0)
    return <Button variant="success" onClick={() => {
        setCount(count + 1);
        console.log(count)
    }} >{count}</Button>
}
export default TestButtons;