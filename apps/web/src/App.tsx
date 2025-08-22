import { useQuery } from "@tanstack/react-query"
import axios from 'axios'
import { Button } from "./components/ui/button"

interface Event{
  id: string,
  title: string,
}

const getEvents = () => {
  return axios.get<Event[]>("http://localhost:8080/events").then(res => res.data)
}

function App() {
  const {data: events, isLoading} = useQuery({ queryKey: ['events'], queryFn: getEvents })

  return (
    <div className="flex flex-col justify-center items-center">
      EVENT MANGER
      {isLoading && <span>Loading... </span>}
      <Button>Ready, SET, GO...</Button>
      <div>
        {events?.map(evnt => {
          return (<div key={evnt.id}>{evnt.title}</div>)
        })}
      </div>
    </div>
  )
}

export default App
