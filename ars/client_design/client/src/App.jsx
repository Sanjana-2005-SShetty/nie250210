
//import Counter from "./Counter"
import FlightCreate from "./flights/FlightCreate"
import FlightEdit from "./flights/FlightEdit"
import FlightList from "./flights/FlightList"
//import FullName from "./FullName"
//import Pageheader from "./header/Pageheader"
import { BrowserRouter, Route, Routes } from 'react-router-dom'


function App() {

  return (
    <>
    
    <div>
      <BrowserRouter>
      <Routes>
        <Route  path="" element={<FlightList/>}/>
        <Route  path="/flights/list" element={<FlightList/>}/>
        <Route  path="/flights/create" element={<FlightCreate/>}/>
        <Route  path="flights/edit/:id" element={<FlightEdit/>}/>
        </Routes>
        </BrowserRouter>
        
    </div>
    
    
    
    
     
    </>
  )
}

export default App