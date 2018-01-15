import * as d3 from 'd3'
import merge from 'lodash/merge'

/**
 * SimpleGraph Options
 *
 * Chart Options:
 * {
 *      chart: The element to use when selecting, see d3 selector
 *      width: The chart width
 *      height: The chart height
 *      margin: { Object of margins for the chart
 *          top: Top margin of the chart
 *          right: Right margin of the chart
 *          bottom: Bottom margin of the chart
 *          left: Left margin of the chart
 *      }
 *      yPrefix: if yPrefix is used, then y value will be formatted with the prefix, see d3 tickFormat,
 *      onDataHover: is a function that will be called whenever you hover over a dot on the chart
 *      scaleX: is a scale value used for scaling the viewbox on x axis, see CSS transform scale
 *      scaleY: is a scale value used for scaling the viewbox on y axis, see CSS transform scale
 * }
 *
 * Chart DataSet:
 * {
 *      dataSets: [
 *          {
 *              label: The label that will be used for this dataset,
 *              styleLine: Function used for styling the line, first parameter is the default style object, return value is a style object
 *              styleDot: Function used for styling the dots, first parameter is the default style object, return value is a style object,
 *              data: [
 *                  {
 *                      y: y axis value, see d3 line data
 *                      x: x axis value, see d3 line data
 *                  }
 *              ]
 *          }
 *      ],
 *      options: {
 *          xFormat: The format to provide for tickFormat on x axis, see d3 tickFormat
 *      }
 * }
 */

let SimpleGraph = function (elem, options = {}) {
    let self = this

    // Default data
    this.data = []
    this.dataOptions = {}

    // Default options
    this._options = {
        chart: elem,
        width: elem.offsetWidth || 1000,
        height: elem.offsetHeight || 400,
        margin: {top: 50, right: 50, bottom: 50, left: 50},
        yPrefix: null,
        onDataHover: null,
        scaleX: 0.95,
        scaleY: 0.75,
        dotHoverScale: 1.2
    }

    // Apply provided options
    this.options(options)

    let scale = this._options.scaleX + ', ' + this._options.scaleY
    let transform = this._options.margin.left + ', ' + this._options.margin.top

    // 1. Add the SVG to the page and employ #2
    this.svg = d3.select(this._options.chart)
        .append('svg')
        .attr('width', '100%')
        .attr('height', '100%')
        .attr('viewBox', '0 0 ' + this._options.width + ' ' + this._options.height)
        .attr('preserveAspectRatio', 'xMinYMin')
        .append('g')
        .attr('transform', 'scale(' + scale + ') translate(' + transform + ')')

    // Define the div for the tooltip
    this.tooltip = d3.select('body').append('div')
        .attr('class', 'tooltip')
        .style('opacity', 0)

    this.tooltip.append('i')
        .attr('class', 'fa fa-times close-tooltip pull-right')
        .on('mouseup', function () {
            console.log(self.tooltip)
            self.tooltip.attr('show', false)
                .transition()
                .duration(500)
                .style('opacity', 0)
        })

    this.tooltip.append('span') // tooltip body
        .attr('class', 'tooltip-body')

    this.redraw()
}

//
// SimpleGraph methods
//

/**
 * options - Options is the method for changing the default options
 * @param {ChartOptions} opts See Chart Options
 */
SimpleGraph.prototype.options = function (opts) {
    // Merge existing options with new options
    this._options = merge(this._options, opts)
}

/**
 * update - Method for updating data and options
 * @param {ChartDataset} data See Chart DataSet
 * @param {ChartOptions} opts See Chart Options
 */
SimpleGraph.prototype.update = function (data, opts = null) {
    if (opts) this.options(opts)
    this.data = data.dataSets || []
    this.dataOptions = data.options
    this.redraw()
}

/**
 * redraw - will update the chart graph
 */
SimpleGraph.prototype.redraw = function () {
    let self = this

    // The number of datapoints
    let n = 0
    let yRange = [-1, -1]
    let xRange = [-1, -1]

    for (let i = 0; i < this.data.length; i++) {
        if (this.data[i].data.length > n) n = this.data[i].data.length
        // Get highest and lowest numbers on x and y
        let yExtent = d3.extent(this.data[i].data, (d, i) => d.y || i)
        let xExtent = d3.extent(this.data[i].data, (d, i) => d.x || i)

        // Compare yRange and yExtent so yRange always have lowest and highest number.
        if (yRange[0] === -1 || yRange[0] > yExtent[0]) yRange[0] = yExtent[0]
        if (yRange[1] === -1 || yRange[1] < yExtent[1]) yRange[1] = yExtent[1]

        // Compare xRange and xExtent so xRange always have lowest and highest number.
        if (xRange[0] === -1 || xRange[0] > xExtent[0]) xRange[0] = xExtent[0]
        if (xRange[1] === -1 || xRange[1] < xExtent[1]) xRange[1] = xExtent[1]
    }

    // Modify the range to give some margin on the range
    yRange[0] = yRange[0] / 1.2
    yRange[1] = yRange[1] * 1.2

    // 5. X scale will use the index of our data
    let xScale = d3.scaleLinear()
        .domain(xRange) // input
        .range([0, this._options.width]) // output

    // 6. Y scale will use the randomly generate number
    let yScale = d3.scaleLinear()
        .domain(yRange) // input
        .range([this._options.height, 0]) // output

    for (let i = 0; i < this.data.length; i++) {
        // 7. d3's line generator
        let line = d3.line()
            .x(function (d, i) { return xScale(d.x || i) }) // set the x values for the line generator
            .y(function (d) { return yScale(d.y) }) // set the y values for the line generator
            .curve(d3.curveMonotoneX) // apply smoothing to the line

        // 3. Call the x axis in a group tag
        let xAxis = d3.axisBottom(xScale)

        if (this.dataOptions.xFormat) {
            xAxis = xAxis.tickFormat(this.dataOptions.xFormat)
        }

        this.svg.append('g')
            .attr('class', 'x axis')
            .attr('transform', 'translate(0,' + this._options.height + ')')
            .style('font-size', '0.9em')
            .call(xAxis) // Create an axis component with d3.axisBottom
            // Rotate x axis text
            .selectAll('text')
            .attr('y', 0)
            .attr('x', 9)
            .attr('dy', '1em')
            .attr('dx', '-1.6em')
            .attr('transform', function (d) {
                return 'rotate(-65)'
            })
            .style('text-anchor', 'end')

        // 4. Call the y axis in a group tag
        let yAxis = d3.axisLeft(yScale)
        if (this._options.yPrefix !== null) {
            yAxis = yAxis.tickFormat(d3.format(this._options.yPrefix))
        } else {
            yAxis = yAxis.tickFormat(d3.format('.3s'))
        }

        this.svg.append('g')
            .attr('class', 'y axis')
            .style('font-size', '0.9em')
            .call(yAxis) // Create an axis component with d3.axisLeft

        // console.log(data[i])
        // 9. Append the path, bind the data, and call the line generator
        this.svg.append('path')
            .datum(this.data[i].data) // 10. Binds data to the line
            .attr('class', 'line') // Assign a class for styling
            .attr('d', line) // 11. Calls the line generator
            .attr('style', styleFunc(
                {'stroke': '#ffab00', 'fill': 'none', 'stroke-width': '3px'},
                self.data[i].styleLine
            ))

        // 12. Appends a circle for each datapoint
        this.svg.selectAll()
            .data(this.data[i].data)
            .enter().append('circle') // Uses the enter().append() method
            .attr('class', 'dot') // Assign a class for styling
            .attr('cx', function (d, i) { return xScale(d.x || i) })
            .attr('cy', function (d) { return yScale(d.y) })
            .attr('r', 5)
            .attr('style', styleFunc(
                {'stroke': '#fff', 'fill': '#ffab00'},
                self.data[i].styleDot
            ))
            .on('mouseover', function (d, i) {
                // Hover over effect
                let radius = d3.select(this).attr('r')
                d3.select(this).attr('r', Math.min(radius * self._options.dotHoverScale))

                if (self.tooltip.attr('show') === true) return
                self.tooltip.transition()
                    .duration(200)
                    .style('opacity', 0.9)
                    .style('left', (d3.event.pageX + 7) + 'px')
                    .style('top', (d3.event.pageY - 125) + 'px')

                let html = d.y
                if (self._options.onDataHover) {
                    html = self._options.onDataHover(d, i)
                }

                self.tooltip.select('span')
                    .html(html)
            })
            .on('mouseout', function (d) {
                // Hover over effect
                let radius = d3.select(this).attr('r')
                d3.select(this).attr('r', Math.min(radius / self._options.dotHoverScale))

                if (self.tooltip.attr('show')) return
                self.tooltip.transition()
                    .duration(500)
                    .style('opacity', 0)
            })
            .on('mouseup', function (d) {
                self.tooltip.attr('show', true)
            })
    }
}

export default SimpleGraph

/**
 * styleFunc - Internal function when handling styling
 * @param {*} def The default style options
 * @param {*} func The function to call for styling
 */
function styleFunc (def, func) {
    return function (d) {
        let style = def
        if (typeof func !== 'undefined') style = func(style)
        let styleString = ''
        for (let key in style) {
            styleString += key + ': ' + style[key] + '; '
        }
        return styleString
    }
}
