import React, { useEffect, useState } from 'react'
import { ApiClient } from '../services/apiClient'

export const AssignmentsPage: React.FC = () => {
  const [list, setList] = useState<any[]>([])

  useEffect(() => {
    ApiClient.get('/assignments').then((data:any) => setList(data)).catch(()=>setList([]))
  }, [])

  return (
    <div>
      <h2>Assignments</h2>
      <ul>
        {list.map(a => <li key={a.id}>{a.title} - due {a.due_date}</li>)}
      </ul>
    </div>
  )
}
