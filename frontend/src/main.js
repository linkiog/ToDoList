import { AddTask, DeleteTask, GetTasks, MarkTaskCompleted } from '../wailsjs/go/backend/App.js';


document.addEventListener('DOMContentLoaded', async () => {
    const addButton = document.getElementById('addButton');
    const taskInput = document.getElementById('taskInput');
    const dueDateInput = document.getElementById('dueDate');
    const priorityInput = document.getElementById('priority');
    const taskList = document.getElementById('taskList');

    if (!addButton || !taskInput || !dueDateInput || !priorityInput || !taskList) {
        console.error('Element not found');
        return;
    }

    const loadTasks = async () => {
        try {
            const tasks = await GetTasks();
            console.log('Loaded tasks:', tasks); // Отладка
            if (tasks.length === 0) {
                taskList.innerHTML = '<li>No tasks available</li>';
            } else {
                taskList.innerHTML = tasks.map(task => `
                    <li class="${task.completed ? 'completed' : ''}">
                        ${task.name} - ${new Date(task.due_date).toLocaleString()} - Priority: ${task.priority}
                        <button class="deleteButton" onclick="deleteTask(${task.id})">Delete</button>
                        <button onclick="markAsCompleted(${task.id})">${task.completed ? 'Undo' : 'Complete'}</button>
                    </li>
                `).join('');
            }
        } catch (error) {
            console.error('Error loading tasks:', error);
        }
    };

    addButton.addEventListener('click', async () => {
        const taskName = taskInput.value.trim();
        const dueDateValue = dueDateInput.value;
        const priority = parseInt(priorityInput.value, 10);

        if (taskName && dueDateValue && !isNaN(priority) && priority >= 1 && priority <= 5) {
            try {
                const dueDate = new Date(dueDateValue).toISOString();
                await AddTask(taskName, dueDate, priority);
                taskInput.value = '';
                dueDateInput.value = '';
                priorityInput.value = '';
                loadTasks();
            } catch (error) {
                console.error('Error adding task:', error);
            }
        } else {
            alert('Please enter valid task details');
        }
    });

    window.deleteTask = async (id) => {
        console.log('Attempting to delete task with ID:', id);
        if (confirm('Are you sure you want to delete this task?')) {
            try {
                await DeleteTask(id);
                console.log('Task deleted successfully');
                loadTasks();
            } catch (error) {
                console.error('Error deleting task:', error);
            }
        } else {
            console.log('Task deletion cancelled');
        }
    };

    window.markAsCompleted = async (id) => {
        try {
            await MarkTaskCompleted(id);
            loadTasks();
        } catch (error) {
            console.error('Error marking task as completed:', error);
        }
    };

    loadTasks();
});
